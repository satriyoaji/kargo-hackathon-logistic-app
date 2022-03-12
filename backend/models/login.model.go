package models

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email" validate:"required,email,unique"`
}
