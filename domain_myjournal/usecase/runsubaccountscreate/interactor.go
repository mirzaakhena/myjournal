package runsubaccountscreate

import (
	"context"
	"fmt"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/shared/infrastructure/database"
	"strings"
)

//go:generate mockery --name Outport -output mocks/

type runSubAccountsCreateInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runSubAccountsCreateInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runSubAccountsCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	subAccountObjs := make([]*entity.SubAccount, 0)

	parentAccountMapUnique := map[entity.AccountCode]any{}
	parentAccountIDList := make([]entity.AccountID, 0)
	for _, account := range req.SubAccounts {

		if _, exist := parentAccountMapUnique[account.ParentAccountCode]; exist {
			continue
		}

		parentAccountMapUnique[account.ParentAccountCode] = nil
		parentAccountIDList = append(parentAccountIDList, entity.NewAccountID(req.WalletId, account.ParentAccountCode))

	}

	p := database.NewDefaultParam().
		Filter("wallet_id", req.WalletId).
		Filter("_id", map[string]any{"$in": parentAccountIDList})

	var parentAccountMap map[entity.AccountCode]entity.Account
	_, err := r.outport.FindAccounts(ctx).GetAllEachItem(p, func(result entity.Account) {
		parentAccountMap[result.Code] = result
	})
	if err != nil {
		return nil, err
	}

	if len(parentAccountMap) == 0 {
		return nil, fmt.Errorf("parent account is Empty")
	}

	for _, account := range req.SubAccounts {

		parentAccount := parentAccountMap[account.ParentAccountCode]

		if account.Code == "" {
			account.Code = entity.SubAccountCode(strings.ToUpper(strings.ReplaceAll(account.Name, " ", "_")))
		}

		subAccountObjs = append(subAccountObjs, &entity.SubAccount{
			ID:            entity.NewSubAccountID(parentAccount.ID, account.Code),
			Code:          account.Code,
			Name:          account.Name,
			ParentAccount: parentAccount,
		})
	}

	err = r.outport.SaveSubAccounts(ctx).InsertMany(subAccountObjs...)
	if err != nil {
		return nil, err
	}

	return res, nil
}
