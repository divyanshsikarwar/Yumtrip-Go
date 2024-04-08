package core

import (
	"context"
	"time"
	"yumtrip/models"
	"yumtrip/utils"
)

//Sends the OTP to the user's email address, by generating a random code and saving it in the database

func SendEmailOtp(context context.Context, email string) error {
	Authentication := models.Authentication{}
	Authentication.Email = email
	Authentication.Code = utils.GenerateRandomCode()
	Authentication.Time = time.Now()
	err := Authentication.create()
	if err != nil {
		return err
	}
	err = utils.SendOtpEmail(email, Authentication.Code)
	if err != nil {
		return err
	}
	return nil
}

func VerifyEmailOtp(context context.Context, otp string) error {
	Authentication := models.Authentication{}
	err := Authentication.get(otp)
	if err != nil {
		return err
	}
	err = Authentication.delete()
	return err
}