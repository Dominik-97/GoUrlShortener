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
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/", handlerFunctions.Slash).Methods("GET")
	muxRouter.HandleFunc("/api/store", handlerFunctions.Store).Methods("POST")
	muxRouter.HandleFunc("/api/getlinks", handlerFunctions.GetLinks).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", muxRouter))
}

func main() {
	fmt.Println("Running the API.")
	// When connection is needed, open connection, execute, and the close connection
	database.ConnectToTheDatabase("configuration/application.yaml")
	// Run the API requestHandler function
	requestHandler()
}
