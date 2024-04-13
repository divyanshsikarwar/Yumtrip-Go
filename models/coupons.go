package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Coupon struct {
	ID            string `json:"_id" bson:"_id"`
	Code          string `json:"code" bson:"code"`
	Discount      string `json:"discount" bson:"discount"` // percentage
	IsActive      bool   `json:"isActive" bson:"isActive"`
	TimesUsed     int    `json:"timesUsed" bson:"timesUsed"`
	MaxUses       int    `json:"maxUses" bson:"maxUses"`
	AvailableFrom string `json:"availableFrom" bson:"availableFrom"`
	Expiration    string `json:"expiration" bson:"expiration"`
}

func (c *Coupon) GetID() string {
	return c.ID
}

func (c *Coupon) Create(ctx context.Context) error {
	_, err := collection.InsertOne(ctx, c)
	if err != nil {
		return err
	}
	return nil
}

func GetCoupons(ctx context.Context) ([]Coupon, error) {
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var coupons []Coupon
	for cursor.Next(ctx) {
		var coupon Coupon
		if err := cursor.Decode(&coupon); err != nil {
			return nil, err
		}
		coupons = append(coupons, coupon)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return coupons, nil
}
