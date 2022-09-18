package runjournalcreate

import (
	"context"
	"fmt"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/domain_myjournal/model/repository"
	"myjournal/shared/infrastructure/database"
	"time"
)

//go:generate mockery --name Outport -output mocks/

type runJournalCreateInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runJournalCreateInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runJournalCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	//balancesList := make([]*entity.SubAccountBalance, 0)

	journalId := entity.NewJournalID(req.WalletId, req.UserId, req.Date)

	subAccountCodes := r.getSubAccountCodes(req.Transactions)

	subAccountObjMap := map[entity.SubAccountCode]entity.SubAccount{}

	p := database.NewDefaultParam().
		Filter("parent_account.wallet_id", req.WalletId).
		Filter("code", map[string]any{"$in": subAccountCodes})

	_, err := r.outport.FindSubAccounts(ctx).GetAllEachItem(p, func(result entity.SubAccount) {
		subAccountObjMap[result.Code] = result
	})
	if err != nil {
		return nil, err
	}

	if len(subAccountObjMap) == 0 {
		return nil, fmt.Errorf("subaccount is empty")
	}

	balanceBySubAccountCodeMap, err := r.outport.FindLastSubAccountBalances(ctx, repository.FindLastSubAccountBalancesRequest{
		WalletID:        req.WalletId,
		SubAccountCodes: subAccountCodes,
	})
	if err != nil {
		return nil, err
	}

	balances, err := r.getBalance(req.Transactions, req.Date, balanceBySubAccountCodeMap, subAccountObjMap, journalId)
	if err != nil {
		return nil, err
	}

	journalObj := entity.Journal{
		ID:          journalId,
		Date:        req.Date,
		WalletID:    req.WalletId,
		UserID:      req.UserId,
		Description: req.Description,
		Balances:    balances,
	}

	err = journalObj.Validate()
	if err != nil {
		return nil, err
	}

	//balancesList = append(balancesList, balances...)

	err = r.outport.SaveJournal(ctx).InsertOrUpdate(journalObj)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveSubAccountBalances(ctx).InsertMany(balances...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r runJournalCreateInteractor) getBalance(
	subAccountBalances []Transaction,
	date time.Time,
	balanceBySubAccountCodeMap map[entity.SubAccountCode]entity.Money,
	subAccountMap map[entity.SubAccountCode]entity.SubAccount,
	journalId entity.JournalID,
) ([]*entity.SubAccountBalance, error) {

	totalBalancePerDirection := map[entity.BalanceDirection]entity.Money{}

	subAccountBalancesResult := make([]*entity.SubAccountBalance, 0)
	for i, sab := range subAccountBalances {

		subAccount, exist := subAccountMap[sab.SubAccountCode]
		if !exist {
			return nil, fmt.Errorf("SubAccountCode %s is not exist", sab.SubAccountCode)
		}

		subAccountBalanceId := entity.NewSubAccountBalanceID(journalId, subAccount.Code)

		lastBalance, exist := balanceBySubAccountCodeMap[sab.SubAccountCode]
		if !exist {
			lastBalance = 0
		}

		direction := subAccount.GetDirection(sab.Sign)

		amount := sab.Sign.GetNumberSign() * sab.Amount

		subAccountBalancesResult = append(subAccountBalancesResult, &entity.SubAccountBalance{
			ID:         subAccountBalanceId,
			JournalID:  journalId,
			Date:       date,
			Amount:     amount,
			Balance:    lastBalance + amount,
			Sequence:   i + 1,
			Direction:  direction,
			SubAccount: subAccount,
		})

		totalBalancePerDirection[direction] += sab.Amount

	}

	if totalBalancePerDirection[entity.BalanceDirectionDebit] != totalBalancePerDirection[entity.BalanceDirectionCredit] {
		return nil, fmt.Errorf("journal is not balance")
	}

	return subAccountBalancesResult, nil
}

func (r runJournalCreateInteractor) getSubAccountCodes(subAccountBalances []Transaction) []entity.SubAccountCode {
	subAccountCodes := make([]entity.SubAccountCode, 0)
	for _, balance := range subAccountBalances {
		subAccountCodes = append(subAccountCodes, balance.SubAccountCode)
	}
	return subAccountCodes
}
