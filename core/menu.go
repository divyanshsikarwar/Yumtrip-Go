package core

import (
	"context"
	"errors"
	"yumtrip/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetStoreMenu(context context.Context, storeId primitive.ObjectID) (models.Menu, error) {
	query := bson.M{
		"storeID": storeId,
	}
	
	menus, err := models.GetStoreMenusByQuery(context,query)
	if err != nil {
		return models.Menu{}, err
	}
	if len(menus) == 0 {
		return models.Menu{}, errors.New("Menu not found")
	}
	return menus[0], nil
}