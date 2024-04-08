package handlers

import (
	"encoding/json"
	"net/http"
	"yumtrip/models"
)

type Menu struct{}

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
	storeID := r.FormValue("storeID")
	menu, err := models.GetStoreMenu(storeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(menu)
}
