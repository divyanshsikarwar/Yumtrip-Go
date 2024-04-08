package models

type Credentials struct {
	ID       string `json:"_id"`
	Key      string `json:"key"`
	Email    string `json:"email"`
	Password string `json:"password"`
}