package handlerFunctions

import (
	"net/http"
	"fmt"
	"UrlShortener/models"
	"UrlShortener/database"
	"encoding/json"
	"strconv"
)

func GetLinks(w http.ResponseWriter, r *http.Request) {
	bookings := []models.Url{}

	var pageInt int
	var perPageInt int

	if page := r.URL.Query().Get("page"); page == "" {
		pageInt, _ = strconv.Atoi("1")
	} else {
		pageInt, _ = strconv.Atoi(page)
	}

	if perPage := r.URL.Query().Get("perPage"); perPage == "" {
		perPageInt, _ = strconv.Atoi("10")
	} else {
		perPageInt, _ = strconv.Atoi(perPage)
	}

	fmt.Println(pageInt, perPageInt)


	if link := r.URL.Query().Get("link"); link != "" {
		fmt.Printf("%s\n", link)
		m := make(map[string]interface{})
		m["full_url"] = link
		database.Db.Where(m).Find(&bookings).Offset(1).Limit(1)
	} else {
		database.Db.Find(&bookings)
	}

	fmt.Println("Endpoint Hit: returnAllBookings")

	fmt.Printf("%T", bookings)

	json.NewEncoder(w).Encode(bookings)
}