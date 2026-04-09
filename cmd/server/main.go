package main

import (
	"net/http"

	"github.com/jpendzich/jpd3v-website/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", handlers.HomeHandler)

	http.ListenAndServe("localhost:8000", mux)
}

