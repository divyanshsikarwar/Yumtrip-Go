package handlers

import (
	"net/http"
	"yumtrip/core"
	"yumtrip/utils"
)

type EmailVerify struct{}

func (e *EmailVerify) SendOtp (w http.ResponseWriter, r *http.Request) {
	
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

func (e *EmailVerify) VerifyOtp (w http.ResponseWriter, r *http.Request) {
	emailOtp := r.FormValue("emailOtp")
	if !utils.IsValidEmailOtp(emailOtp) {
		http.Error(w, "Invalid email otp", http.StatusBadRequest)
		return
	}

	err := core.VerifyEmailOtp(r.Context(), emailOtp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Send the success response
	w.WriteHeader(http.StatusOK)
}