package entity

type UserID string

func (u UserID) String() string {
	return string(u)
}

type User struct {
	ID   UserID `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
