package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Authentication struct {
	ID    string `json:"_id"`
	Email string `json:"email"`
	Code  string `json:"code"`
	Time  time.Time `json:"time"`
}

func (a *Authentication) Create() error {
	return nil
}
func (a *Authentication) Delete() error {
	return nil
}

func GetAuthenticationDocByQuery(query bson.M) (Authentication, error) {
	return Authentication{}, nil
}