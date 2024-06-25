package main

import (
	"e-commerce-back/internal/auth"
	"e-commerce-back/pkg/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database := db.DBConnection()

	authHandler := auth.NewHandler(database)

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()

	authHandler.RegisterRoutes(apiRouter)

	log.Fatal(http.ListenAndServe(":8080", r))
}
