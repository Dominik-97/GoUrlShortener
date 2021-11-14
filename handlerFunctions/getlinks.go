package handlerFunctions

import (
	"net/http"
	"fmt"
)

func GetLinks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World.")
}