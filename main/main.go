package main

import (
	"net/http"

	"github.com/cyoa"
)

func main() {
	handler := cyoa.MainHandler{}
	http.ListenAndServe(":8080", handler)
}
