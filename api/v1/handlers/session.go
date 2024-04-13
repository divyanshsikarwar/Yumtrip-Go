package handlers

import (
	"encoding/json"
	"net/http"
	"yumtrip/api/v1/models"
	"yumtrip/core"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//To be depricated, moved to Oauth2.0
type Session struct{}

func (s *Session) ValidateSession(w http.ResponseWriter, r *http.Request) {
	var sessionId primitive.ObjectID
	sessionIdHex := r.FormValue("sessionId")
	sessionId, err := primitive.ObjectIDFromHex(sessionIdHex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	valid, userid, err := core.ValidateSession(r.Context(), sessionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !valid {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	//Update exisiting session
	err = core.CreateUpdateSession(r.Context() , sessionId ,userid)
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

	valid, sessionId, userId, err := creds.Validate(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !valid {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	err = core.CreateUpdateSession(r.Context(), sessionId, userId)
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
	valid, _, err := core.ValidateSession(r.Context(),sessionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if valid {
		err = core.InValidateSession(r.Context(), sessionId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}