package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yumtrip/core"
	"yumtrip/models"
)

type Store struct{}


func (s *Store) CreateStore(w http.ResponseWriter, r *http.Request) {
	//Parse the request body
	var store models.Store
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Create the store
	err = store.Create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Send the success response
	w.WriteHeader(http.StatusCreated)
}

func (s *Store) GetStores(w http.ResponseWriter, r *http.Request) {
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

	stores, err := core.GetStores(r.Context(), skip, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stores)
}

func (s *Store) GetStore(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	store, err := core.GetStore(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(store)
}

func (s *Store) UpdateStore(w http.ResponseWriter, r *http.Request) {
	var store models.Store
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = store.Update()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Store) DeleteStore(w http.ResponseWriter, r *http.Request) {
	var store models.Store
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = store.Delete()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}