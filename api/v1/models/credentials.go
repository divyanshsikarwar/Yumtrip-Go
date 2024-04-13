package models

import (
	"yumtrip/core"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Credentials struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (c *Credentials) Validate() (bool, primitive.ObjectID, error) {
	email := c.Email
	password := c.Password
	user, err := core.GetUserByEmail(email)
	if err != nil {
		return false, primitive.NilObjectID, err
	}
	if user.ID == primitive.NilObjectID {
		return false, primitive.NilObjectID, nil
	}
	userId := user.ID
	originalPassword, err := core.GetPasswordHashByUserId(userId)
	if err != nil {
		return false, primitive.NilObjectID, err
	}
	if originalPassword == password {
		sessionId := core.GetSessionIdByUserId(userId)
		return true, sessionId, nil
	}
	return false, primitive.NilObjectID, nil
}