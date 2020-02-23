package model

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Telephone string `json:"telephone"`
	Password string `json:"password"`
}
