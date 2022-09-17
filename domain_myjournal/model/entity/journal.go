package entity

import (
	"fmt"
	"time"
)

type JournalID string

func NewJournalID(walletId WalletID, userId UserID, now time.Time) JournalID {
	return JournalID(fmt.Sprintf("%s_%s_%s", walletId, userId, now.Format("060102150405")))
}

type Journal struct {
	ID          JournalID            `json:"id" bson:"_id"`
	Date        time.Time            `json:"date" bson:"date"`
	WalletID    WalletID             `json:"walletId" bson:"wallet_id"  index:"-1"`
	UserID      UserID               `json:"userId" bson:"user_id"`
	Description string               `json:"description" bson:"description"`
	Balances    []*SubAccountBalance `json:"balances" bson:"balances"`
}

func (j *Journal) Validate() error {
	return nil
}
