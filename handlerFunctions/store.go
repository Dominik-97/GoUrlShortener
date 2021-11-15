package handlerFunctions

import (
	"net/http"
	"fmt"
	"time"
	"UrlShortener/models"
	"encoding/json"
	"UrlShortener/database"
)

func Store(w http.ResponseWriter, r *http.Request) {
	//Read and store the body, as it can not be read multiple times
	decoder := json.NewDecoder(r.Body)

	//Convert the byte array to a map
	var m map[string]interface{}
	err := decoder.Decode(&m)
    if err != nil {
        panic(err)
    }

	//Check if the required keys are present
	if _, ok := m["fullUrl"]; !ok {
		json.NewEncoder(w).Encode(map[string]string{"error": "fullUrl key is missing from the request."})
		return
	}
	if _, ok := m["shortenedUrl"]; !ok {
		json.NewEncoder(w).Encode(map[string]string{"error": "shortenedUrl key is missing from the request."})
		return
	}

	// Add the date value to the request
	m["date"] = GetCurrentDate()

	// Marshal the request data again, so it can be later Unmarshaled to the UrlObject struct
	// TODO - Refactor this logic
	newData, err := json.Marshal(m)
	if err != nil {
		panic("Problem")
	}

	// Unmarshal the data to the UrlObject struct
    var UrlObject models.Url
    json.Unmarshal(newData, &UrlObject)
    
	// Check if database connection is established and connect if it does not

	// Create the record in the database
    database.Db.Create(&UrlObject)

	// Return the created object
    fmt.Println("Endpoint Hit: Creating New Url")
    json.NewEncoder(w).Encode(UrlObject)
}

func GetCurrentDate() string {
	t := time.Now()
	currentDate := fmt.Sprintf("%d/%d/%d", t.Day(), int(t.Month()), t.Year())
	return currentDate
}
