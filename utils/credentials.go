package utils

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// IsValidEmail is a function that checks if the email is valid
func IsValidEmail(email string) bool {
	//Regular expression for email validation
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

//Generate hash of the password
func GetPasswordHash(password string) (string, error) {
	//Hash the password with the salt and return the hash
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}