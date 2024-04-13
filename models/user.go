package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the model for the user
type User struct {
	ID       primitive.ObjectID `json:"id"`
	StoreID  primitive.ObjectID `json:"storeID"`
	Email    string `json:"email"`
	ImageUrl string `json:"imageUrl"`
	Role     Role `json:"role"`
}

func (u *User) GetID() string {
	return u.ID.Hex()
}

//create user in mongo db
func (u *User) Create() error {
	// u.create()
	return nil
}

func (u *User) Update() error {
	// u.update()
	return nil
}

func GetUserByQuery(query bson.M) (User, error) {
	// get
	return User{}, nil
}