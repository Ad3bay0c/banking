package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/create", Create).Methods("GET")
	router.HandleFunc("/customer/{id:[0-9]+}", Greet)

	log.Print("Server Started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
