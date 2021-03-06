// リスト3.10
package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// ハンドラー関数をチェインするパターン
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("ハンドラ関数が呼び出されました - " + name)
		h(w, r)
	}
}

// ハンドラーをチェインするパターン。

// HelloHandler implements http.Handler
type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("ハンドラが呼び出されました - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// some code to make sure the user is authorized
		h.ServeHTTP(w, r)
	})
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/method", log(hello))
	hello := HelloHandler{}
	http.Handle("/handler", protect(log(hello)))
	server.ListenAndServe()
}
