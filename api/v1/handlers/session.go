package handlers

import (
	"encoding/json"
	"net/http"
	"yumtrip/core"
	"yumtrip/api/v1/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct{}

func (s *Session) ValidateSession(w http.ResponseWriter, r *http.Request) {
	var sessionId primitive.ObjectID
	sessionIdHex := r.FormValue("sessionId")
	sessionId, err := primitive.ObjectIDFromHex(sessionIdHex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	valid, err := core.ValidateSession(sessionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !valid {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	//Update exisiting session
	err = core.CreateUpdateSession(sessionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Session) Login(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	valid, sessionId, err := creds.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !valid {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	err = core.CreateUpdateSession(sessionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Session) Logout(w http.ResponseWriter, r *http.Request) {
	var sessionId primitive.ObjectID
	sessionIdHex := r.FormValue("sessionId")
	sessionId, err := primitive.ObjectIDFromHex(sessionIdHex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	valid, err := core.ValidateSession(sessionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if valid {
		err = core.InValidateSession(sessionId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}