package main

import (
	"net/http"
	"time"

	"github.com/gblcarvalho/go-expert-rate-limiter/pkg/ratelimiter"
	"github.com/go-chi/chi/v5"
	// "github.com/redis/go-redis/v9"
)

func main() {
	store := ratelimiter.NewMemoryStore()
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })
	// store := ratelimiter.NewRedisStore(rdb)
	opts := ratelimiter.RateLimiterMiddlewareOpts{
		IPMaxRequests:    10,
		TokenMaxRequests: 100,
		TimeWindow:       time.Millisecond * 1000,
	}
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
