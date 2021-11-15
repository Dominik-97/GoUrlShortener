package database

import (
	"fmt"
	"log"
	"UrlShortener/configReader"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"UrlShortener/database/models"
)

//interface
//type db struct {
//	var db *gorm.DB
//}

var Db *gorm.DB
var err error

func LoadConfig(configLocation string) configReader.Conf {
	var config configReader.Conf
	config.GetConf(configLocation)

	return config
}

// Implement connect to the database function

func ConnectToTheDatabase(configLocation string) {
	config := LoadConfig(configLocation)
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", 
		config.DatabaseUser, 
		config.DatabasePassword,
		config.DatabaseURL,
		config.DatabasePort,
		config.DatabaseName)

	Db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic("Could not connect to the database.")
	}else{ 
		log.Println("Connection Established")
	}

	// db.AutoMigrate(&database.Url{})
}

// Implement AutoMigrateFunc

func AutoMigrate() {
	Db.AutoMigrate(&database.Url{})
}

// Implement the close connection to the database function

func CloseConnection(db *gorm.DB) error {
	return db.Close()
}
