package handlers

import (
	"encoding/json"
	"net/http"
	"yumtrip/core"
	"yumtrip/models"
)

type User struct{}

type UserWithCreds struct {
	user models.User
	password string
}


func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	//Parse the request body
	var userAndCreds UserWithCreds
	err := json.NewDecoder(r.Body).Decode(&userAndCreds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := userAndCreds.user
	password := userAndCreds.password

	//Create the user
	err = core.CreateUserAndCredentials(user, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Send the success response
	w.WriteHeader(http.StatusCreated)
}

//Mostly will be used to update role of the user
func (u *User) UpdateUser(w http.ResponseWriter, r *http.Request) {
	//Parse the request body
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Update the user
	err = user.Update()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Send the success response
	w.WriteHeader(http.StatusOK)
}