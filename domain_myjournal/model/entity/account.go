package entity

import "fmt"

type AccountSide string

const AccountSideActiva = AccountSide("ACTIVA")
const AccountSidePassiva = AccountSide("PASSIVA")

type AccountId string

func NewAccountId(id WalletId, code AccountCode) AccountId {
	return AccountId(fmt.Sprintf("%s_%s", id, code))
}

type AccountLevel int

type AccountCode string

type Account struct {
	Id              AccountId    `json:"id" bson:"_id"`
	WalletId        WalletId     `json:"walletId" bson:"wallet_id"`
	Code            AccountCode  `json:"code" bson:"code" index:"1"`
	Name            string       `json:"name" bson:"name"`
	Level           AccountLevel `json:"level" bson:"level"`
	Side            AccountSide  `json:"side" bson:"side"`
	ParentAccountID AccountId    `json:"parentAccountId,omitempty" bson:"parent_account_id"`
}
