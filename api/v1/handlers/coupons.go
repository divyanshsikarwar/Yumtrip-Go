package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yumtrip/core"
	"yumtrip/models"
)

type Coupon struct{}

func (c *Coupon) CreateCoupon(w http.ResponseWriter, r *http.Request) {
	//Parse the request body
	var coupon models.Coupon
	err := json.NewDecoder(r.Body).Decode(&coupon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Create the coupon
	err = coupon.Create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Send the success response
	w.WriteHeader(http.StatusCreated)
}

func (c *Coupon) GetCoupons(w http.ResponseWriter, r *http.Request) {
	skipStr := r.FormValue("skip")
	limitStr := r.FormValue("limit")
	
	skip,err := strconv.Atoi(skipStr)
	if err != nil {
		http.Error(w, "Invalid skip value", http.StatusBadRequest)
		return
	}
	limit ,err := strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, "Invalid limit value", http.StatusBadRequest)
		return
	}

	coupons, err := core.GetCoupons(r.Context(), skip, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(coupons)
}
