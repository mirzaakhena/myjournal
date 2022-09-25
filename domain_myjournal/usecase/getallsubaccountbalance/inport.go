package getallsubaccountbalance

import (
	"context"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/domain_myjournal/model/repository"
	"time"

	"myjournal/shared/usecase"
)

type Inport usecase.Inport[context.Context, InportRequest, InportResponse]

// InportRequest is request payload to run the usecase
type InportRequest struct {
	repository.FindAllSubAccountBalanceRequest
	Page int64 `form:"page,omitempty,default=1"`
	Size int64 `form:"size,omitempty,default=30"`
}

//type SubAccountBalanceItem struct {
//	SubAccountBalance entity.SubAccountBalance `json:"sub_account_balance"`
//	Journal           entity.Journal           `json:"journal"`
//}

// InportResponse is response payload after running the usecase
type InportResponse struct {
	Count int64
	Items []TheItem
}

type TheItem struct {
	ID         entity.SubAccountBalanceID `json:"id" bson:"_id"`
	Journal    Journal                    `json:"journal" bson:"journal"`
	SubAccount SubAccount                 `json:"subAccount" bson:"sub_account"`
	Date       time.Time                  `json:"date" bson:"date"`
	Amount     string                     `json:"amount" bson:"amount"`
	Balance    string                     `json:"balance" bson:"balance"`
	Sequence   int                        `json:"sequence" bson:"sequence"`
	Direction  entity.BalanceDirection    `json:"direction" bson:"direction"`
}

type Journal struct {
	ID          entity.JournalID `json:"id" bson:"_id"`
	UserID      entity.UserID    `json:"userId" bson:"user_id"`
	Description string           `json:"description" bson:"description"`
}

type SubAccount struct {
	ID            entity.SubAccountID   `json:"id" bson:"_id"`
	Code          entity.SubAccountCode `json:"code" bson:"code" index:"1"`
	Name          string                `json:"name" bson:"name"`
	ParentAccount Account               `json:"parentAccount" bson:"parent_account"`
}

type Account struct {
	ID       entity.AccountID    `json:"id" bson:"_id"`
	WalletID entity.WalletID     `json:"walletId" bson:"wallet_id"  index:"-1"`
	Code     entity.AccountCode  `json:"code" bson:"code" index:"1"`
	Name     string              `json:"name" bson:"name"`
	Level    entity.AccountLevel `json:"level" bson:"level"`
	Side     entity.AccountSide  `json:"side" bson:"side"`
}
