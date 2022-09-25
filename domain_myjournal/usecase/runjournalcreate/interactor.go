package runjournalcreate

import (
	"context"
	"fmt"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/domain_myjournal/model/repository"
	"sort"
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

	journalId := entity.NewJournalID(req.WalletId, req.UserId, req.Date)

	journalObj := entity.Journal{
		ID:          journalId,
		Date:        req.Date,
		WalletID:    req.WalletId,
		UserID:      req.UserId,
		Description: req.Description,
	}

	err := r.outport.SaveJournal(ctx, &journalObj)
	if err != nil {
		return nil, err
	}

	subAccountCodes := make([]entity.SubAccountCode, 0)
	for _, balance := range req.Transactions {
		subAccountCodes = append(subAccountCodes, balance.SubAccountCode)
	}

	subAccountObjMap, err := r.outport.FindSubAccounts(ctx, repository.FindSubAccountsRequest{
		WalletID:        journalObj.WalletID,
		SubAccountCodes: subAccountCodes,
	})
	if err != nil {
		return nil, err
	}

	if len(subAccountObjMap) == 0 {
		return nil, fmt.Errorf("subaccount is empty")
	}

	balanceBySubAccountCodeMap, err := r.outport.FindLastSubAccountBalances(ctx, repository.FindLastSubAccountBalancesRequest{
		WalletID:        journalObj.WalletID,
		SubAccountCodes: subAccountCodes,
	})
	if err != nil {
		return nil, err
	}

	totalBalancePerDirection := map[entity.BalanceDirection]entity.Money{}

	subAccountBalancesResult := make([]*entity.SubAccountBalance, 0)
	for _, sab := range req.Transactions {

		subAccount, exist := subAccountObjMap[sab.SubAccountCode]
		if !exist {
			return nil, fmt.Errorf("SubAccountCode %s is not exist", sab.SubAccountCode)
		}

		subAccountBalanceId := entity.NewSubAccountBalanceID(journalObj.ID, subAccount.Code)

		lastBalance, exist := balanceBySubAccountCodeMap[sab.SubAccountCode]
		if !exist {
			lastBalance = 0
		}

		direction := subAccount.GetDirection(sab.Sign)

		amount := sab.Sign.GetNumberSign() * sab.Amount

		subAccountBalancesResult = append(subAccountBalancesResult, &entity.SubAccountBalance{
			ID:         subAccountBalanceId,
			JournalID:  journalObj.ID,
			UserID:     journalObj.UserID,
			Date:       journalObj.Date,
			Amount:     amount,
			Balance:    lastBalance + amount,
			Sequence:   0,
			Direction:  direction,
			SubAccount: subAccount,
		})

		totalBalancePerDirection[direction] += sab.Amount

	}

	if totalBalancePerDirection[entity.BalanceDirectionDebit] != totalBalancePerDirection[entity.BalanceDirectionCredit] {
		return nil, fmt.Errorf("journal is not balance")
	}

	// arrange the sequence of DEBIT and CREDIT (debit first before credit)
	sort.Slice(subAccountBalancesResult, func(i, j int) bool {
		return subAccountBalancesResult[i].Direction > subAccountBalancesResult[j].Direction
	})

	for i, x := range subAccountBalancesResult {
		x.Sequence = i + 1
	}

	err = r.outport.SaveSubAccountBalances(ctx, subAccountBalancesResult)
	if err != nil {
		return nil, err
	}

	return res, nil
}
