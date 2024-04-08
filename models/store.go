package models

type Store struct {
	ID          string   `json:"_id"`
	Key         string   `json:"key"`
	Email       string   `json:"email"`
	Name        string   `json:"name"`
	Phone        string   `json:"phone"`
	Address     string   `json:"address"`
	City        string   `json:"city"`
	Logo        string   `json:"logo"`
	Image       string   `json:"image"`
	Rating      string   `json:"rating"`
	Reviews     []string `json:"reviews"`
}