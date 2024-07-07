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

	items.RegisterEndpoints(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}