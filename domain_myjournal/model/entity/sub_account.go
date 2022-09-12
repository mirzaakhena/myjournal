package entity

import "fmt"

type SubAccountId string

type SubAccountCode string

type SubAccount struct {
	Id            SubAccountId   `json:"id" bson:"_id"`
	Code          SubAccountCode `json:"code" bson:"code" index:"1"`
	Name          string         `json:"name" bson:"name"`
	ParentAccount Account        `json:"parentAccount" bson:"parent_account"`
}

func NewSubAccountId(id AccountId, code SubAccountCode) SubAccountId {
	return SubAccountId(fmt.Sprintf("%s_%s", id, code))
}

func (b SubAccount) GetDirection(sign SubAccountBalanceSign) BalanceDirection {

	if (sign == SubAccountBalanceSignPlus && b.ParentAccount.Side == AccountSideActiva) ||
		(sign == SubAccountBalanceSignMinus && b.ParentAccount.Side == AccountSidePassiva) {
		return BalanceDirectionDebit
	}

	if (sign == SubAccountBalanceSignPlus && b.ParentAccount.Side == AccountSidePassiva) ||
		(sign == SubAccountBalanceSignMinus && b.ParentAccount.Side == AccountSideActiva) {
		return BalanceDirectionCredit
	}

	return ""
}
