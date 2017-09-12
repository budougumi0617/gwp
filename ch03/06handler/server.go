package main

import (
	"fmt"
	"net/http"
)

// MyHandler is inplements Handler.
type MyHandler struct{}

// ServeHTTP puts strings.
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
