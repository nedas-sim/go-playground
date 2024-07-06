package main

import (
	"crud-app/database"
	"crud-app/handlers/items"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initialized := database.InitializeDatabase()
	if !initialized {
		return
	}

	log.Println("Database connection successful")

	r := mux.NewRouter()

	r.HandleFunc("/api/items", items.Get).Methods("GET")
	r.HandleFunc("/api/items", items.Create).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}