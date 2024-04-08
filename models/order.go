package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Order struct {
	ID          string `json:"_id"`
	Key 	 string `json:"key"`
	Time 	  time.Time `json:"time"`
	Items 	 []string `json:"items"`  // array of item IDs
	Email 	 string `json:"email"`
	Phone 	 string `json:"phone"`
	TableNo int `json:"tableNo"`
	Status string `json:"status"` // pending, preparing, ready, delivered
	TotalPrice float64 `json:"totalPrice"`
	Paid bool `json:"paid"`
	TotalPaid float64 `json:"totalPaid"`
	CouponCode string `json:"couponCode"`
}

func (o *Order) GetID() string {
	return o.ID
}

func (o *Order) GetKey() string {
	return o.Key
}

func (o *Order) CreateOrder() error {
	// Create Order in MongoDB o.createOrder()
	return nil
}

func GetOrders() ([]Order, error) {
	// Get Orders from MongoDB
	return nil, nil
}

func GetOrdersByQuery(query bson.M) ([]Order, error) {
	// Get Orders from MongoDB by query
	return nil, nil
}