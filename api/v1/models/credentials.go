package models

import (
	"context"
	"yumtrip/core"
	"yumtrip/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Credentials struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (c *Credentials) Validate(context context.Context) (bool, primitive.ObjectID, primitive.ObjectID, error) {
	email := c.Email
	password := c.Password
	user, err := core.GetUserByEmail(context, email)
	if err != nil {
		return false, primitive.NilObjectID,primitive.NilObjectID, err
	}
	if user.ID == primitive.NilObjectID {
		return false, primitive.NilObjectID,primitive.NilObjectID, nil
	}
	userId := user.ID
	originalPassword, err := core.GetPasswordHashByUserId( context,userId)
	if err != nil {
		return false, primitive.NilObjectID,primitive.NilObjectID, err
	}
	isPasswordCorrect := utils.CheckPassword(password, originalPassword)
	if isPasswordCorrect {
		sessionId := core.GetSessionIdByUserId(context,userId)
		return true, sessionId, user.ID, nil
	}
	return false, primitive.NilObjectID,primitive.NilObjectID, nil
}