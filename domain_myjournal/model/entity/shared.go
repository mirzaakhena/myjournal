package entity

type Money float64

type BalanceDirection string

const BalanceDirectionDebit = BalanceDirection("DEBIT")
const BalanceDirectionCredit = BalanceDirection("CREDIT")

type SubAccountBalanceSign string

const SubAccountBalanceSignPlus = SubAccountBalanceSign("+")
const SubAccountBalanceSignMinus = SubAccountBalanceSign("-")

func (s SubAccountBalanceSign) GetNumberSign() Money {
	return map[SubAccountBalanceSign]Money{
		SubAccountBalanceSignPlus:  Money(1),
		SubAccountBalanceSignMinus: Money(-1),
	}[s]
}

type HasIDEntity[T ~string] interface {
	GetID() T
}

type BaseEntity[T ~string] struct {
	Id T `json:"id" bson:"_id"`
}

func (b BaseEntity[T]) GetID() T {
	return b.Id
}
