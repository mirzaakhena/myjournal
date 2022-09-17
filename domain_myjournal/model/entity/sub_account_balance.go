package entity

import (
	"fmt"
	"time"
)

type SubAccountBalanceID string

func NewSubAccountBalanceID(journalId JournalID, subAccountCode SubAccountCode) SubAccountBalanceID {
	return SubAccountBalanceID(fmt.Sprintf("%s_%s", journalId, subAccountCode))
}

type SubAccountBalance struct {
	ID         SubAccountBalanceID `json:"id" bson:"_id"`
	JournalID  JournalID           `json:"journalId" bson:"journal_id" index:"-1"`
	SubAccount SubAccount          `json:"subAccount" bson:"sub_account"`
	Date       time.Time           `json:"date" bson:"date"`
	Amount     Money               `json:"amount" bson:"amount"`
	Balance    Money               `json:"balance" bson:"balance"`
	Sequence   int                 `json:"sequence" bson:"sequence"`
	Direction  BalanceDirection    `json:"direction" bson:"direction"`
}
