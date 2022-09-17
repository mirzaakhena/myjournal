package runsubaccountscreate

import (
	"context"
	"fmt"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/domain_myjournal/model/repository"
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
	parentAccountList := make([]entity.AccountID, 0)
	for _, account := range req.SubAccounts {

		if _, exist := parentAccountMapUnique[account.ParentAccountCode]; exist {
			continue
		}

		parentAccountMapUnique[account.ParentAccountCode] = nil
		parentAccountList = append(parentAccountList, entity.NewAccountID(req.WalletId, account.ParentAccountCode))

	}

	parentAccountMap, err := r.outport.FindAccounts(ctx, repository.FindAccountsRequest{
		WalletID:   req.WalletId,
		AccountIds: parentAccountList,
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

	err = r.outport.SaveSubAccounts(ctx, subAccountObjs)
	if err != nil {
		return nil, err
	}

	return res, nil
}
