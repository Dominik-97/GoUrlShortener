package handlerFunctions

import (
	"net/http"
	"fmt"
	"time"
	"io/ioutil"
	"UrlShortener/database/models"
	"encoding/json"
	"UrlShortener/db"
)

func Store(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
    // return the string response containing the request body
    reqBody, _ := ioutil.ReadAll(r.Body)
    var booking database.Url
    json.Unmarshal(reqBody, &booking)
    db.Create(&booking) 
    fmt.Println("Endpoint Hit: Creating New Booking")
    json.NewEncoder(w).Encode(booking)
}

func GetCurrentDate() string {
	t := time.Now()
	currentDate := fmt.Sprintf("%d/%d/%d", t.Day(), int(t.Month()), t.Year())
	return currentDate
}