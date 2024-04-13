package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Menu struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	StoreID         primitive.ObjectID `json:"storeID" bson:"storeID"`
	Items           []Item             `json:"items" bson:"items"`
	BackgroundImage string             `json:"backgroundImage" bson:"backgroundImage"`
}

func (m *Menu) GetID() primitive.ObjectID {
	return m.ID
}

func UpsertMenu(ctx context.Context, menu Menu) error {
	opts := options.Replace().SetUpsert(true)
	_, err := collection.ReplaceOne(ctx, bson.M{"_id": menu.ID}, menu, opts)
	if err != nil {
		return err
	}
	return nil
}

func GetStoreMenusByQuery(ctx context.Context, query bson.M) ([]Menu, error) {
	cursor, err := collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var menus []Menu
	for cursor.Next(ctx) {
		var menu Menu
		if err := cursor.Decode(&menu); err != nil {
			return nil, err
		}
		menus = append(menus, menu)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return menus, nil
}
