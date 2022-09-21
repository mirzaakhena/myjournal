package entity

type WalletID string

func (u WalletID) String() string {
	return string(u)
}

type Wallet struct {
	ID   WalletID `json:"id" bson:"id"`
	Name string   `json:"name" bson:"name"`
}
