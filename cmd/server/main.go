package main

import (
	"net/http"
	"time"

	"github.com/gblcarvalho/go-expert-rate-limiter/pkg/ratelimiter"
	"github.com/go-chi/chi/v5"
)

func main() {
	opts := ratelimiter.RateLimiterMiddlewareOpts{
		IPMaxRequests:    10,
		TokenMaxRequests: 100,
		TimeWindow:       time.Millisecond * 1000,
	}
	store := ratelimiter.NewMemoryStore()
	rateLimiterMiddleware, err := ratelimiter.NewRateLimiterMiddleware(store, opts)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(rateLimiterMiddleware.Handler)

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
	})

	http.ListenAndServe(":8080", r)
}
