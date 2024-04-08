package handlers

import (
	"encoding/json"
	"net/http"
	"yumtrip/models"
)

type Session struct{}

func (s *Session) Login(w http.ResponseWriter, r *http.Request) {
	var session models.Session
	err := json.NewDecoder(r.Body).Decode(&session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Validate the login
	err = session.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Create the session
	err = session.CreateUpdateSession()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Send the success response
	w.WriteHeader(http.StatusCreated)
}

func (s *Session) Logout(w http.ResponseWriter, r *http.Request) {
	var session models.Session
	err := json.NewDecoder(r.Body).Decode(&session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Delete the session
	err = session.Delete()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Send the success response
	w.WriteHeader(http.StatusOK)
}