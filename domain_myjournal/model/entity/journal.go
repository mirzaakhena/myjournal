package entity

import (
	"fmt"
	"time"
)

type JournalId string

func NewJournalId(walletId WalletId, userId UserId, now time.Time) JournalId {
	return JournalId(fmt.Sprintf("%s_%s_%s", walletId, userId, now.Format("060102150405")))
}

type Journal struct {
	Id          JournalId            `json:"id" bson:"_id"`
	Date        time.Time            `json:"date" bson:"date"`
	WalletId    WalletId             `json:"walletId" bson:"wallet_id"  index:"-1"`
	UserId      UserId               `json:"userId" bson:"user_id"`
	Description string               `json:"description" bson:"description"`
	Balances    []*SubAccountBalance `json:"balances" bson:"balances"`
}

func (j *Journal) Validate() error {
	return nil
}
