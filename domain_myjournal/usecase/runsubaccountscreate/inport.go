package runsubaccountscreate

import (
	"context"
	"myjournal/domain_myjournal/model/entity"

	"myjournal/shared/usecase"
)

type Inport usecase.Inport[context.Context, InportRequest, InportResponse]

type SubAccount struct {
	ParentAccountCode entity.AccountCode    `json:"parentAccountCode"`
	Code              entity.SubAccountCode `json:"code,omitempty"`
	Name              string                `json:"name"`
}

type InportRequest struct {
	WalletId    entity.WalletId `json:"walletId"`
	SubAccounts []SubAccount    `json:"subAccounts"`
}

type InportResponse struct {
}
