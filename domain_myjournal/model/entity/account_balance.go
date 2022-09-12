package entity

import (
	"time"
)

type AccountBalanceId string

type AccountBalance struct {
	Id      AccountBalanceId `json:"id" bson:"_id"`
	Account Account          `json:"account" bson:"account"`
	Date    time.Time        `json:"date" bson:"date"`
	Balance Money            `json:"balance" bson:"balance"`
}
