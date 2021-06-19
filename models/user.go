package models

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (u User) TableName() string {
	return "user"
}
