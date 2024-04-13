package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Credentials struct {
	ID       primitive.ObjectID `json:"_id"`
	UserId   primitive.ObjectID `json:"userId"`
	Password string `json:"password"`
}

func (c *Credentials) GetID() string {
	return c.ID.Hex()
}
func (c *Credentials) GetPasswordHash() string {
	return c.Password
}

func (c *Credentials) Create() error {
	// c.create()
	return nil
}

func GetPasswordHashByUserId(userId primitive.ObjectID) (string, error) {
	return "", nil
}