package utils

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

//string validation that string is alpha numeric
func IsValidEmailOtp(emailOtp string) bool {
	//Regular expression for email otp validation
	if len(emailOtp) != 6 {
		return false
	}
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(emailOtp)
}

func GenerateRandomCode() (string, error) {
	bycryptPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bycryptPassword), nil
}