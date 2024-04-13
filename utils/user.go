package utils

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserFromContext(ctx context.Context) (primitive.ObjectID, error) {
	user, ok := ctx.Value("user").(primitive.ObjectID)
	if !ok {
		return user, errors.New("User not found in context")
	}
	return user, nil
}