package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Role is the role of the user
type Role struct {
	ID   primitive.ObjectID `json:"id"`
	StoreID primitive.ObjectID `json:"storeID"`
	Name string `json:"name"`
	Permissions []Permission `json:"permissions"`
}

// Permission is the permission of the role
type Permission string
//Permission could be any of the following : View_Analytics, Manage_Orders, Manage_Menu, Manage_Coupons, Manage_Users

func (r *Role) GetID() string {
	return r.ID.Hex()
}

func (r *Role) Create() error {
	// r.create()
	return nil
}

func (r *Role) Update() error {
	// r.update()
	return nil
}

func GetRolesByQuery(query bson.M) ([]Role, error) {
	// Get Roles from MongoDB
	return nil, nil
}