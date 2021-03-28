package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", ServiceHandler)
	http.ListenAndServe(":7777", nil)
}
