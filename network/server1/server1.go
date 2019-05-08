package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// each request calls handler
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
}
