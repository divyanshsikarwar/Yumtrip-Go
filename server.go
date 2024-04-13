package main

import (
	"log"
	"net/http"
	apiv1 "yumtrip/api/v1"
	cronworkers "yumtrip/workers/cronWorkers"

	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()
	v1Router := router.PathPrefix("/api/v1").Subrouter()
	v1Router.Handle("/", apiv1.Router())

	log.Println("Server is starting...")

	cronworkers.Init()

	// Start the server
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
