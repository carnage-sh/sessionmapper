package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", DisplayHandler)
	http.ListenAndServe(":5000", nil)
}
