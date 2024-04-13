package core

import (
	"yumtrip/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPasswordHashByUserId(userId primitive.ObjectID) (string, error) {
	return models.GetPasswordHashByUserId(userId)
}