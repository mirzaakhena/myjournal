package entity

import (
	"time"
)

type AccountBalanceID string

type AccountBalance struct {
	ID      AccountBalanceID `json:"id" bson:"_id"`
	Account Account          `json:"account" bson:"account"`
	Date    time.Time        `json:"date" bson:"date"`
	Balance Money            `json:"balance" bson:"balance"`
}
