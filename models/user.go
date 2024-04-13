// models.go
package models

import (
	"context"
	"errors"
	"yumtrip/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection *mongo.Collection
)

func init() {
	// Set collection
	collection = database.GetCollection("users")
}

// User is the model for the user
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StoreID  primitive.ObjectID `json:"storeID" bson:"storeID"`
	Email    string             `json:"email" bson:"email"`
	ImageUrl string             `json:"imageUrl" bson:"imageUrl"`
	Role     Role               `json:"role" bson:"role"`
}

func (u *User) GetID() string {
	return u.ID.Hex()
}

// Create user in MongoDB
func (u *User) Create(ctx context.Context) error {
	_, err := collection.InsertOne(ctx, u)
	if err != nil {
		return err
	}
	return nil
}

// Update user in MongoDB
func (u *User) Update(ctx context.Context) error {
	filter := bson.M{"_id": u.ID}
	update := bson.M{"$set": u}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByQuery retrieves a user from MongoDB based on query
func GetUserByQuery(ctx context.Context,query bson.M) (User, error) {
	var user User
	err := collection.FindOne(ctx, query).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return User{}, errors.New("user not found")
		}
		return User{}, err
	}
	return user, nil
}
