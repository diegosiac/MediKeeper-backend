package main

import (
	"e-commerce-back/internal/auth"
	"e-commerce-back/pkg/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()

	r := mux.NewRouter()
	r.HandleFunc("/login", auth.LoginHandler).Methods("POST")
	r.HandleFunc("/register", auth.RegisterHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
