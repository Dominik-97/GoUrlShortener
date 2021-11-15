package handlerFunctions

import (
	"net/http"
	"fmt"
	"time"
	//"io/ioutil"
	"UrlShortener/models"
	"encoding/json"
	"UrlShortener/database"
)

func Store(w http.ResponseWriter, r *http.Request) {
	//Read and store the body, as it can not be read multiple times
    //reqBody, _ := ioutil.ReadAll(r.Body)
	decoder := json.NewDecoder(r.Body)

	//Convert the byte array to a map
	var m map[string]interface{}
	//err := json.Unmarshal(reqBody, &m)
	//if err != nil {
	//	panic("There was a problem converting the request.")
	//}

	err := decoder.Decode(&m)
    if err != nil {
        panic(err)
    }
    //log.Println(t.Test)
	fmt.Println(m)

	if val, ok := m["fullUrl"]; ok {
		fmt.Println(val, "Yes")
	}else{
		fmt.Println("No")
	}
	//if val, ok := m["foo"]; ok {
	//	fmt.Println(val, "Yes")
	//}else{
	//	fmt.Println("No")
	//	fmt.Fprintf(w, "Hello World.")
	//	return
	//}
	m["date"] = GetCurrentDate()
	newData, err := json.Marshal(m)
	if err != nil {
		panic("Problem")
	}
	//fmt.Println(reqBody)
	//fmt.Println(r.Body)
	//fmt.Printf("%s", newData)
    var UrlObject models.Url
    json.Unmarshal(newData, &UrlObject)
    // add currentDate
    database.Db.Create(&UrlObject) 
    //fmt.Println("Endpoint Hit: Creating New Url")
    json.NewEncoder(w).Encode(UrlObject)
	//fmt.Fprintf(w, "Hello World.")
}

func GetCurrentDate() string {
	t := time.Now()
	currentDate := fmt.Sprintf("%d/%d/%d", t.Day(), int(t.Month()), t.Year())
	return currentDate
}