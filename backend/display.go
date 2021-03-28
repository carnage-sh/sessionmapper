package main

import (
	"fmt"
	"net/http"
)

func DisplayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Headers\n")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}
}
