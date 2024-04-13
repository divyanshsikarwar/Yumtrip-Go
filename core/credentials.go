package core

import (
	"context"
	"yumtrip/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPasswordHashByUserId(context context.Context ,userId primitive.ObjectID) (string, error) {
	return models.GetPasswordHashByUserID(context,userId)
}