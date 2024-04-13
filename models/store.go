// store.go
package models

import (
	"context"
	"yumtrip/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	storeCollection *mongo.Collection
)

func init() {
	// Set collection
	storeCollection = database.GetCollection("stores")
}

// Store is the model for the store
type Store struct {
	ID       string   `json:"_id"`
	Key      string   `json:"key"`
	Email    string   `json:"email"`
	Name     string   `json:"name"`
	Phone    string   `json:"phone"`
	Address  string   `json:"address"`
	City     string   `json:"city"`
	Logo     string   `json:"logo"`
	Image    string   `json:"image"`
	Rating   string   `json:"rating"`
	Reviews  []string `json:"reviews"`
}

func (s *Store) GetID() string {
	return s.ID
}

// Create a store in MongoDB
func (s *Store) Create(ctx context.Context) error {
	_, err := storeCollection.InsertOne(ctx, s)
	if err != nil {
		return err
	}
	return nil
}

// Update a store in MongoDB
func (s *Store) Update(ctx context.Context) error {
	filter := bson.M{"_id": s.ID}
	update := bson.M{"$set": s}
	_, err := storeCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete a store from MongoDB
func (s *Store) Delete(ctx context.Context) error {
	filter := bson.M{"_id": s.ID}
	_, err := storeCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

// GetStoresByQuery retrieves stores from MongoDB based on query
func GetStoresByQuery(ctx context.Context, query bson.M) ([]Store, error) {
	cursor, err := storeCollection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stores []Store
	for cursor.Next(ctx) {
		var store Store
		if err := cursor.Decode(&store); err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return stores, nil
}
