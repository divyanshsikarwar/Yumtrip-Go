package handlers

import (
	"net/http"
	"yumtrip/core"
	"yumtrip/utils"
)

type Authentication struct{}

func (e *Authentication) SendOtp (w http.ResponseWriter, r *http.Request) {
	
	email := r.FormValue("email")
	if !utils.IsValidEmail(email) {
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}
	err := core.SendEmailOtp(r.Context(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Send the success response
	w.WriteHeader(http.StatusCreated)
}

func (e *Authentication) VerifyOtp (w http.ResponseWriter, r *http.Request) {
	emailOtp := r.FormValue("emailOtp")
	if !utils.IsValidEmailOtp(emailOtp) {
		http.Error(w, "Invalid email otp", http.StatusBadRequest)
		return
	}
	email := r.Context().Value("email").(string)
	
	isVerified := core.VerifyEmailOtp(r.Context(),emailOtp, email)
	if !isVerified {
		http.Error(w, "Invalid email otp", http.StatusBadRequest)
		return
	}

	//Send the success response
	w.WriteHeader(http.StatusOK)
}