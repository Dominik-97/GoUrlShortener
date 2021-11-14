package handlerFunctions

import (
	"net/http"
	"fmt"
	"time"
)

func Store(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World.")
}

func GetCurrentDate() string {
	t := time.Now()
	currentDate := fmt.Sprintf("%d/%d/%d", t.Day(), int(t.Month()), t.Year())
	return currentDate
}