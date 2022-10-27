package models

type User struct {
	Id        int    `json:"_id"`
	FirstName string `json:"firstname"`
	Age       int    `json:"age"`
}
