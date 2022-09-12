package entity

type UserId string

func (u UserId) String() string {
	return string(u)
}

type User struct {
	Id   UserId `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
