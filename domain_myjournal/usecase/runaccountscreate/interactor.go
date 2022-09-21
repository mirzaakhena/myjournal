package runaccountscreate

import (
	"context"
	"fmt"
	"myjournal/domain_myjournal/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type runAccountsCreateInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &runAccountsCreateInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *runAccountsCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	if len(req.RootAccounts) == 0 {
		return nil, fmt.Errorf("accounts must > 0")
	}

	accountObjs := make([]*entity.Account, 0)
	accountObjs = r.traceAccounts(ctx, accountObjs, req.WalletId, 1, req.RootAccounts, "", "")

	err := r.outport.SaveAccounts(ctx, accountObjs)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *runAccountsCreateInteractor) traceAccounts(ctx context.Context, accountObjs []*entity.Account, walletId entity.WalletID, level entity.AccountLevel, childAccounts []Account, parentAccountId entity.AccountID, parentSide entity.AccountSide) []*entity.Account {

	for _, account := range childAccounts {

		accountId := entity.NewAccountID(walletId, account.Code)

		if account.Side == "" {
			account.Side = parentSide
		}

		accountObjs = append(accountObjs, &entity.Account{
			ID:              accountId,
			WalletId:        walletId,
			Code:            account.Code,
			Name:            account.Name,
			Level:           level,
			Side:            account.Side,
			ParentAccountID: parentAccountId,
		})

		accountObjs = r.traceAccounts(ctx, accountObjs, walletId, level+1, account.Accounts, accountId, account.Side)

	}

	return accountObjs
}
