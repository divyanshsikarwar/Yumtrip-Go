package models

type Item struct {
	ID          string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Stock       int `json:"stock"`
	ImageUrl       string `json:"imageUrl"`
	IsAvailable bool `json:"isAvailable"`
}