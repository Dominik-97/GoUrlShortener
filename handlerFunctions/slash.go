package handlerFunctions

import (
	"net/http"
	"fmt"
)

func Slash(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World.")
}