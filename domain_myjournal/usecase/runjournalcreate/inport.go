package runjournalcreate

import (
	"context"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/shared/usecase"
	"time"
)

type Inport usecase.Inport[context.Context, InportRequest, InportResponse]

// InportRequest is request payload to run the usecase
type InportRequest struct {
	WalletId     entity.WalletId `json:"-"`
	UserId       entity.UserId   `json:"-"`
	Date         time.Time       `json:"-"`
	Description  string          `json:"description,omitempty"`
	Transactions []Transaction   `json:"transactions"`
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
}

type Transaction struct {
	Sign           entity.SubAccountBalanceSign `json:"sign"`
	SubAccountCode entity.SubAccountCode        `json:"subAccountCode"`
	Amount         entity.Money                 `json:"amount"`
}
