package core

import (
	"context"
	"errors"
	"log"
	"yumtrip/models"
	"yumtrip/utils"
)

func CreateUserAndCredentials(context context.Context,user models.User, password string) error {
	//TODO : Make this atomic
	passwordHex, err := utils.GetPasswordHash(password)
	if err != nil {
		return errors.New("Error hashing the password err : "+ err.Error())
	}

	err = user.Create(context)
	if err != nil {
		return errors.New("Error creating user err : "+ err.Error())
	}

	creds := models.Credentials{
		UserID: user.ID,
		Password: passwordHex,
	}

	err = creds.Create(context)
	if err != nil {
		log.Println("[SEVERE] Error creating user credentials UserId : "+ user.ID.Hex(), err) 
	}
	return nil
}

func GetUserByEmail(context context.Context ,email string) (models.User, error) {
	query := map[string]interface{}{
		"email": email,
	}
	user, err := models.GetUserByQuery(context ,query)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}