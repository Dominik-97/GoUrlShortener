package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"UrlShortener/handlerFunctions"
	"UrlShortener/configReader"
)

func requestHandler() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/", handlerFunctions.Slash).Methods("GET")
	muxRouter.HandleFunc("/api/store", handlerFunctions.Store).Methods("GET")
	muxRouter.HandleFunc("/api/getlinks", handlerFunctions.GetLinks).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", muxRouter))
}

func main() {
	fmt.Println("Ha")

	var config configReader.Conf
	config.GetConf("configuration/application.yaml")

	fmt.Println(config)
	requestHandler()
}