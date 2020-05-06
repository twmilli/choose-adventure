package cyoa

import (
	"fmt"
	"net/http"
)

// MainHandler Main HTTP Handler interface
type MainHandler struct{}

func (h MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
