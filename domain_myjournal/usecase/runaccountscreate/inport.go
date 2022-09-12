package runaccountscreate

import (
	"context"
	"myjournal/domain_myjournal/model/entity"

	"myjournal/shared/usecase"
)

type Inport usecase.Inport[context.Context, InportRequest, InportResponse]

type Account struct {
	Code     entity.AccountCode `json:"code"`
	Name     string             `json:"name"`
	Side     entity.AccountSide `json:"side,omitempty"`
	Accounts []Account          `json:"accounts,omitempty"`
}

type InportRequest struct {
	WalletId     entity.WalletId `json:"walletId"`
	RootAccounts []Account       `json:"accounts"`
}

type InportResponse struct {
}
