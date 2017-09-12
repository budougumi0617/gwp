package main

import (
	"fmt"
	"net/http"
)

// HelloHandler implements Handler interface.
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// WorldHandler implements Handler interface.
type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
		// Handlerを指定しないと、DefaultServeMuxをハンドラとして利用する。
	}

	http.Handle("/hello", &hello) // hellハンドラをDefaultServeMuxに付加する。
	http.Handle("/world", &world)

	server.ListenAndServe()
}
