package core

import (
	"context"
	"time"
	"yumtrip/models"
	"go.mongodb.org/mongo-driver/bson"
)


func GetOrders(context context.Context, skip, limit int) ([]models.Order, error) {
	orders,err := models.GetOrders(context)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func GetNewInactiveOrders(context context.Context) ([]models.Order, error) {
	//Write mongo query to get all orders that are 3hrs old, active and in pending or delivered stage
	query := bson.M{
		"status": bson.M{
			"$in": []string{"pending", "delivered"},
		},
		"time": bson.M{
			"$lt": time.Now().Add(-3 * time.Hour),
		},
	}
	orders, err := models.GetOrdersByQuery(context, query)
	return orders, err

}