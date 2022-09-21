package entity

import "fmt"

type AccountSide string

const AccountSideActiva = AccountSide("ACTIVA")
const AccountSidePassiva = AccountSide("PASSIVA")

type AccountID string

func NewAccountID(id WalletID, code AccountCode) AccountID {
	return AccountID(fmt.Sprintf("%s_%s", id, code))
}

type AccountLevel int

type AccountCode string

type Account struct {
	ID              AccountID    `json:"id" bson:"_id"`
	WalletId        WalletID     `json:"walletId" bson:"wallet_id"`
	Code            AccountCode  `json:"code" bson:"code" index:"1"`
	Name            string       `json:"name" bson:"name"`
	Level           AccountLevel `json:"level" bson:"level"`
	Side            AccountSide  `json:"side" bson:"side"`
	ParentAccountID AccountID    `json:"parentAccountId,omitempty" bson:"parent_account_id"`
}
