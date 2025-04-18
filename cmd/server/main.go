package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func helloWorldMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello World")
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := chi.NewRouter()

	r.Use(helloWorldMiddleware)

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
	})

	http.ListenAndServe(":8080", r)
}
