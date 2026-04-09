package handlers

import (
	"net/http"
	"os"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	filebytes, _ := os.ReadFile("templates/pages/home.html")
	writer.Write(filebytes)
}
