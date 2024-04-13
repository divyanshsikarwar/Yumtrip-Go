package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yumtrip/core"
	"yumtrip/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct{}

func (rl *Role) CreateRole(w http.ResponseWriter, r *http.Request) {
	//Parse the request body
	var role models.Role
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Create the role
	err = core.CreateRole(role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Send the success response
	w.WriteHeader(http.StatusCreated)
}

func (rl *Role) GetRoles(w http.ResponseWriter, r *http.Request) {
	skipStr := r.FormValue("skip")
	limitStr := r.FormValue("limit")

	skip, err := strconv.Atoi(skipStr)
	if err != nil {
		http.Error(w, "Invalid skip value", http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, "Invalid limit value", http.StatusBadRequest)
		return
	}

	roles, err := core.GetRoles(r.Context(), skip, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(roles)
}

func (rl *Role) GetRole(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	//check if the id is a valid ObjectId
	roleId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid role id", http.StatusBadRequest)
		return
	}

	role, err := core.GetRole(r.Context(), roleId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(role)
}
