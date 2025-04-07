package handlers

import (
	"fmt"
	"net/http"
)

func LongToShort(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	fmt.Fprintf(w, "Welcome to the long to short url!")
}
