package handlerFunctions

import (
	"net/http"
	"fmt"
	"time"
	"io/ioutil"
	//"UrlShortener/models"
	"encoding/json"
	//"UrlShortener/database"
)

func Store(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body
    // check whether json request contains all the required keys
    reqBody, _ := ioutil.ReadAll(r.Body)
	var m map[string]interface{}
	err := json.Unmarshal(reqBody, &m)
	if err != nil {
		panic("Problem")
	}
	if val, ok := m["foo"]; ok {
		fmt.Println(val, "Yes")
	}else{
		fmt.Println("No")
	}
	if val, ok := m["foo"]; ok {
		fmt.Println(val, "Yes")
	}else{
		fmt.Println("No")
		fmt.Fprintf(w, "Hello World.")
		return
	}
	m["new_key"] = "test"
	newData, err := json.Marshal(m)
	if err != nil {
		panic("Problem")
	}
	//fmt.Println(reqBody)
	//fmt.Println(r.Body)
	fmt.Printf("%s", newData)
    //var UrlObject models.Url
    //json.Unmarshal(reqBody, &UrlObject)
    // add currentDate
    //database.Db.Create(&UrlObject) 
    //fmt.Println("Endpoint Hit: Creating New Url")
    //json.NewEncoder(w).Encode(UrlObject)
	fmt.Fprintf(w, "Hello World.")
}

func GetCurrentDate() string {
	t := time.Now()
	currentDate := fmt.Sprintf("%d/%d/%d", t.Day(), int(t.Month()), t.Year())
	return currentDate
}