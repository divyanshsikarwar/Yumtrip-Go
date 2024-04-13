package core

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateSession(sessionId primitive.ObjectID) (bool,error) {
	return true,nil
}

func CreateUpdateSession(sessionId primitive.ObjectID) error {
	return nil
}

func GetSessionIdByUserId(userId primitive.ObjectID) primitive.ObjectID {
	return primitive.NilObjectID
}
func InValidateSession(sessionId primitive.ObjectID) error {
	return nil
}