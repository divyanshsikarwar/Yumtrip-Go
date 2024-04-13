package core

import (
	"errors"
	"log"
	"yumtrip/models"
	"yumtrip/utils"
)

func CreateUserAndCredentials(user models.User, password string) error {
	//TODO : Make this atomic
	passwordHex, err := utils.GetPasswordHash(password)
	if err != nil {
		return errors.New("Error hashing the password err : "+ err.Error())
	}

	err = user.Create()
	if err != nil {
		return errors.New("Error creating user err : "+ err.Error())
	}

	creds := models.Credentials{
		UserId: user.ID,
		Password: passwordHex,
	}

	err = creds.Create()
	if err != nil {
		log.Println("[SEVERE] Error creating user credentials UserId : "+ user.ID.Hex(), err) 
	}
	return nil
}

func GetUserByEmail(email string) (models.User, error) {
	query := map[string]interface{}{
		"email": email,
	}
	user, err := models.GetUserByQuery(query)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}