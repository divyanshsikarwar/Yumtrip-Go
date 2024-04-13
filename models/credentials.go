package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Credentials struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"userId" bson:"userId"`
	Password string             `json:"password" bson:"password"`
}

func (c *Credentials) GetID() string {
	return c.ID.Hex()
}

func (c *Credentials) GetPasswordHash() string {
	return c.Password
}

func (c *Credentials) Create(ctx context.Context) error {
	_, err := collection.InsertOne(ctx, c)
	if err != nil {
		return err
	}
	return nil
}

func GetPasswordHashByUserID(ctx context.Context, userID primitive.ObjectID) (string, error) {
	var credentials Credentials
	err := collection.FindOne(ctx, bson.M{"userId": userID}).Decode(&credentials)
	if err != nil {
		return "", err
	}
	return credentials.Password, nil
}
