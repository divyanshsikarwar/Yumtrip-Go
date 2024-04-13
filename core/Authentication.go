package core

import (
	"context"
	"time"
	"yumtrip/models"
	"yumtrip/utils"

	"go.mongodb.org/mongo-driver/bson"
)

//Sends the OTP to the user's email address, by generating a random code and saving it in the database

func SendEmailOtp(context context.Context, email string) error {
	Authentication := models.Authentication{}
	Authentication.Email = email
	randomCode, err := utils.GenerateRandomCode()
	if err != nil {
		return err
	}
	Authentication.Code = randomCode
	Authentication.Time = time.Now()
	err = Authentication.Create()
	if err != nil {
		return err
	}
	err = utils.SendEmailOtp(email, Authentication.Code)
	if err != nil {
		return err
	}
	return nil
}

func VerifyEmailOtp(context context.Context, otp,email string) bool {
	query := bson.M{"email": email, "code": otp}

	auth, err := models.GetAuthenticationDocByQuery(query)
	if err != nil {
		return false
	}
	auth.Delete()
	return true
}