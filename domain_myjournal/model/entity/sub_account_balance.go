package entity

import (
	"fmt"
	"time"
)

type SubAccountBalanceId string

func NewSubAccountBalanceId(journalId JournalId, subAccountCode SubAccountCode) SubAccountBalanceId {
	return SubAccountBalanceId(fmt.Sprintf("%s_%s", journalId, subAccountCode))
}

type SubAccountBalance struct {
	Id         SubAccountBalanceId `json:"id" bson:"_id"`
	JournalId  JournalId           `json:"journalId" bson:"journal_id" index:"-1"`
	SubAccount SubAccount          `json:"subAccount" bson:"sub_account"`
	Date       time.Time           `json:"date" bson:"date"`
	Amount     Money               `json:"amount" bson:"amount"`
	Balance    Money               `json:"balance" bson:"balance"`
	Sequence   int                 `json:"sequence" bson:"sequence"`
	Direction  BalanceDirection    `json:"direction" bson:"direction"`
}
