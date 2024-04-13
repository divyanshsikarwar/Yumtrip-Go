package apiv1

import (
	"yumtrip/api/v1/handlers"

	"github.com/gorilla/mux"
)

// Router is the main router for the API
func Router() *mux.Router {
	mux := mux.NewRouter()

	// Menu routes
	menu := handlers.Menu{}
	
	mux.HandleFunc("menu", menu.UpdateCreateMenu).Methods("POST")
	mux.HandleFunc("menu", menu.GetStoreMenu).Methods("GET")

	// Coupon routes
	coupon := handlers.Coupon{}
	mux.HandleFunc("coupon", coupon.CreateCoupon).Methods("POST")
	mux.HandleFunc("coupon", coupon.GetCoupons).Methods("GET")

	// Order routes
	order := handlers.Order{}
	mux.HandleFunc("order", order.CreateOrder).Methods("POST")
	mux.HandleFunc("order", order.GetOrders).Methods("GET")

	// Store routes
	store := handlers.Store{}
	mux.HandleFunc("store", store.CreateStore).Methods("POST")
	mux.HandleFunc("store", store.GetStores).Methods("GET")

	// User routes
	user := handlers.User{}
	mux.HandleFunc("user", user.CreateUser).Methods("POST")
	mux.HandleFunc("user", user.UpdateUser).Methods("PUT")
	
	//role routes
	role := handlers.Role{}
	mux.HandleFunc("role", role.CreateRole).Methods("POST")
	mux.HandleFunc("role", role.GetRoles).Methods("GET")
	mux.HandleFunc("role/{role_id}", role.GetRole).Methods("GET")

	// Authentication routes
	auth := handlers.Authentication{}
	mux.HandleFunc("auth", auth.SendOtp).Methods("GET")
	mux.HandleFunc("auth", auth.VerifyOtp).Methods("POST")

	// Authourization routes
	login := handlers.Session{}
	mux.HandleFunc("auth/login", login.Login).Methods("POST")
	mux.HandleFunc("auth/validate", login.ValidateSession).Methods("GET")
	mux.HandleFunc("auth/logout", login.Logout).Methods("DELETE")
	return mux
}