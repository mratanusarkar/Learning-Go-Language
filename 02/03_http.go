package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")

}

// Server function is a basic demo of nr=et/http
func Server() {
	http.HandleFunc("/", indexHandler)

	fmt.Println("server live at: http://localhost:8000/")
	http.ListenAndServe(":8000", nil)
}
