package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/jamesblockk/goRestfulStarterKit/route"
)

func main() {

	router := route.New() // create routes

	// These two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Launch server with CORS validations
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
