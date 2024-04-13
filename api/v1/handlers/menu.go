package handlers

import (
	"encoding/json"
	"net/http"
	"yumtrip/core"
	"yumtrip/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct{}
//TODO : Add Redis Cache for List Apis

func (m *Menu) UpdateCreateMenu(w http.ResponseWriter, r *http.Request) {
	var menu models.Menu
	err := json.NewDecoder(r.Body).Decode(&menu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.UpsertMenu(menu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (m *Menu) GetStoreMenu(w http.ResponseWriter, r *http.Request) {
	storeIDHex := r.FormValue("storeID")
	storeId, err := primitive.ObjectIDFromHex(storeIDHex)
	if err != nil {
		http.Error(w, "Invalid storeID", http.StatusBadRequest)
		return
	}

	menu, err := core.GetStoreMenu(r.Context(), storeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(menu)
}
