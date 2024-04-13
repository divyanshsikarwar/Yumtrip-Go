package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yumtrip/core"
	"yumtrip/models"
)

type Order struct {}

func (o *Order) GetOrders(w http.ResponseWriter, r *http.Request) {
	skipStr := r.FormValue("skip")
	limitStr := r.FormValue("limit")
	//convert to int
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

	orders, err := core.GetOrders(r.Context(), skip, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Send the orders as a response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (o *Order) CreateOrder(w http.ResponseWriter, r *http.Request) {
	//Parse the request body
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Create the order
	err = order.CreateOrder(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Send the success response
	w.WriteHeader(http.StatusCreated)
}