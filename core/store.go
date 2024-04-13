package core

import (
	"context"
	"errors"
	"yumtrip/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetStores(context context.Context, skip, limit int) ([]models.Store, error) {
	query := bson.M{
		"status": "active",
		"skip":   skip,
		"limit":  limit,
	}

	stores, err := models.GetStoresByQuery(query)
	if err != nil {
		return nil, err
	}
	return stores, nil
}

func GetStore(context context.Context, id string) (models.Store, error) {
	query := bson.M{
		"_id":     id,
	}
	stores, err := models.GetStoresByQuery(query)
	if err != nil {
		return models.Store{}, err
	}
	if len(stores)==0 {
		return models.Store{}, errors.New("Store not found")
	}

	return stores[0], nil
}