package main

import (
	"fmt"
	"log"
	"net/http"
)

func ColorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Colors</h1><style>*{background-color: #006400;}</style>")
}

func main() {
	// to add : url functionality
	http.HandleFunc("/color", ColorHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
