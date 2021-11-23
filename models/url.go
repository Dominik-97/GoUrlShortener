package models

type Url struct {
	Id 			 int `json:"id"`
	Date 		 string `json:"date"`
	FullUrl 	 string `json:"fullUrl"`
	ShortenedUrl string `json:"ShortenedUrl"`
	TimesClicked int `json:"timesClicked"`
}

type UrlPaginated struct {
	Urls 	[]Url `json:"urls"`
	Page 	int `json:"page"`
	PerPage int `json:"PerPage"`
}