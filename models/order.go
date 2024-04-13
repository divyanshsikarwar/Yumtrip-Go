package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Key         string             `json:"key" bson:"key"`
	Time        time.Time          `json:"time" bson:"time"`
	Items       []string           `json:"items" bson:"items"`
	Email       string             `json:"email" bson:"email"`
	Phone       string             `json:"phone" bson:"phone"`
	TableNo     int                `json:"tableNo" bson:"tableNo"`
	Status      string             `json:"status" bson:"status"`
	TotalPrice  float64            `json:"totalPrice" bson:"totalPrice"`
	Paid        bool               `json:"paid" bson:"paid"`
	TotalPaid   float64            `json:"totalPaid" bson:"totalPaid"`
	CouponCode  string             `json:"couponCode" bson:"couponCode"`
}

func (o *Order) GetID() string {
	return o.ID.Hex()
}

func (o *Order) CreateOrder(ctx context.Context) error {
	_, err := collection.InsertOne(ctx, o)
	if err != nil {
		return err
	}
	return nil
}

func GetOrders(ctx context.Context) ([]Order, error) {
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []Order
	for cursor.Next(ctx) {
		var order Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrdersByQuery(ctx context.Context, query bson.M) ([]Order, error) {
	cursor, err := collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []Order
	for cursor.Next(ctx) {
		var order Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *Order) UpdateOrder(ctx context.Context) error {
	filter := bson.M{"_id": o.ID}
	update := bson.M{"$set": o}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
