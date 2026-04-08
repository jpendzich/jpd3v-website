package handlers

import (
	"net/http"
	"fmt"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	fmt.Fprint(writer, "<h1>Hello, World!<h1>")
}
