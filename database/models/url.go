package database

type Url struct {
	Id int `json:"id"`
	Date string `json:"date"`
	FullUrl string `json:"fullUrl"`
	ShortenedUrl string `json:"ShortenedUrl"`
}
