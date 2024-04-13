package core

import (
	"context"
	"errors"
	"slices"
	"yumtrip/constants"
	"yumtrip/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateRole (role models.Role) error {
	for _, permission := range role.Permissions {
		if !slices.Contains(constants.AllPermissions, permission) {
			return errors.New("Invalid permission : "+ string(permission))
		}
	}

	err := role.Create()
	if err != nil {
		return err
	}
	return nil
}

func GetRoles(context context.Context, skip, limit int) ([]models.Role, error) {
	query := bson.M{
		"skip":  skip,
		"limit": limit,
	}

	roles, err := models.GetRolesByQuery(query)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func GetRole(context context.Context, id primitive.ObjectID) (models.Role, error) {
	query := bson.M{
		"_id": id,
	}
	roles, err := models.GetRolesByQuery(query)
	if err != nil {
		return models.Role{}, err
	}
	if len(roles) == 0 {
		return models.Role{}, errors.New("Role not found")
	}
	return roles[0], nil
}