package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"UrlShortener/handlerFunctions"
	"UrlShortener/configReader"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"UrlShortener/database/models"
) 

func requestHandler() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/", handlerFunctions.Slash).Methods("GET")
	muxRouter.HandleFunc("/api/store", handlerFunctions.Store).Methods("POST")
	muxRouter.HandleFunc("/api/getlinks", handlerFunctions.GetLinks).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", muxRouter))
}

var db *gorm.DB
var err error

func main() {
	fmt.Println("Ha")

	var config configReader.Conf
	config.GetConf("configuration/application.yaml")

	fmt.Println(config)
	requestHandler()

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", 
		config.DatabaseUser, 
		config.DatabasePassword,
		config.DatabaseURL,
		config.DatabasePort,
		config.DatabaseName)

	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic("Could not connect to the database.")
	}else{ 
		log.Println("Connection Established")
	}

	defer db.Close()

	db.AutoMigrate(&database.Url{})
}