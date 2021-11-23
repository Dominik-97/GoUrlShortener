package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"UrlShortener/handlerFunctions"
	"UrlShortener/database"
) 

func requestHandler() {
	// Create a new router
	muxRouter := mux.NewRouter().StrictSlash(true)

	// Handle routes
	muxRouter.HandleFunc("/", handlerFunctions.Slash).Methods("GET")
	muxRouter.HandleFunc("/api/store", handlerFunctions.Store).Methods("POST")
	muxRouter.HandleFunc("/api/getlinks", handlerFunctions.GetLinks).Methods("GET")

	// Start listening on port 8081
	log.Fatal(http.ListenAndServe(":8081", muxRouter))
}

func main() {
	fmt.Println("Running the API.")

	// Connect to the database using config provided in application.yaml
	database.ConnectToTheDatabase("configuration/application.yaml")

	// Run the API requestHandler function
	requestHandler()
}
