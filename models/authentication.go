package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Authentication struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Email string             `json:"email" bson:"email"`
	Code  string             `json:"code" bson:"code"`
	Time  time.Time          `json:"time" bson:"time"`
}

func (a *Authentication) Create(ctx context.Context) error {
	_, err := collection.InsertOne(ctx, a)
	if err != nil {
		return err
	}
	return nil
}

func (a *Authentication) Delete(ctx context.Context) error {
	filter := bson.M{"_id": a.ID}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func GetAuthenticationDocByQuery(ctx context.Context, query bson.M) (Authentication, error) {
	var auth Authentication
	err := collection.FindOne(ctx, query).Decode(&auth)
	if err != nil {
		return Authentication{}, err
	}
	return auth, nil
}
