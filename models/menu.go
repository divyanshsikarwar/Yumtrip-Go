package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID          primitive.ObjectID `json:"_id"`
	StoreId  primitive.ObjectID `json:"storeID"`
	Items []Item `json:"items"`
	BackgroundImage string `json:"backgroundImage"`
}

func (m *Menu) GetID() primitive.ObjectID {
	return m.ID
}

func UpsertMenu (menu Menu) error {
	return nil
}

func GetStoreMenusByQuery(context context.Context, query bson.M) ([]Menu, error) {
	return []Menu{}, nil
}