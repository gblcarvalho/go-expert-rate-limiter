package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func getIPAddress(r *http.Request) string {
	ips := r.Header.Get("X-Forwarded-For")
	if ips != "" {
		ipsArr := strings.Split(ips, ",")
		return strings.TrimSpace(ipsArr[0])
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func getAPIKEYToken(r *http.Request) string {
	return r.Header.Get("API_KEY")
}

func helloWorldMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Hello World from: %s and Token: %s", getIPAddress(r), getAPIKEYToken(r))
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
