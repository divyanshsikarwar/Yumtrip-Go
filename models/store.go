package models

import "go.mongodb.org/mongo-driver/bson"

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

func (s *Store) GetID() string {
	return s.ID
}

func (s *Store) Create() error {
	// s.create()
	return nil
}

func (s *Store) Update() error {
	// s.update()
	return nil
}

func (s *Store) Delete() error {
	// s.delete()
	return nil
}

func GetStoresByQuery(query bson.M) ([]Store, error) {
	// Get Stores from MongoDB
	return nil, nil
}