package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Role is the role of the user
type Role struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StoreID     primitive.ObjectID `json:"storeID" bson:"storeID"`
	Name        string             `json:"name" bson:"name"`
	Permissions []Permission       `json:"permissions" bson:"permissions"`
}

// Permission is the permission of the role
type Permission string

// Permission values
const (
	ViewAnalytics Permission = "View_Analytics"
	ManageOrders  Permission = "Manage_Orders"
	ManageMenu    Permission = "Manage_Menu"
	ManageCoupons Permission = "Manage_Coupons"
	ManageUsers   Permission = "Manage_Users"
)

func (r *Role) GetID() string {
	return r.ID.Hex()
}

// Create a role in MongoDB
func (r *Role) Create(ctx context.Context) error {
	_, err := collection.InsertOne(ctx, r)
	if err != nil {
		return err
	}
	return nil
}

// Update a role in MongoDB
func (r *Role) Update(ctx context.Context) error {
	filter := bson.M{"_id": r.ID}
	update := bson.M{"$set": r}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

// GetRolesByQuery retrieves roles from MongoDB based on query
func GetRolesByQuery(ctx context.Context, query bson.M) ([]Role, error) {
	cursor, err := collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var roles []Role
	for cursor.Next(ctx) {
		var role Role
		if err := cursor.Decode(&role); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return roles, nil
}
