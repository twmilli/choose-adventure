package main

import (
	"net/http"

	"github.com/cyoa"
)

func main() {
	handler := cyoa.NewMainHandler()
	http.ListenAndServe(":8080", handler)
}
