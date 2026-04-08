package main

import (
	"net/http"

	"github.com/jpendzich/jpd3v-website/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)

	http.ListenAndServe("localhost:8000", mux)
}

