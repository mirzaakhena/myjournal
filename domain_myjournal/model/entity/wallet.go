package entity

type WalletId string

func (u WalletId) String() string {
	return string(u)
}

type Wallet struct {
	Id   WalletId `json:"id" bson:"id"`
	Name string   `json:"name" bson:"name"`
}
