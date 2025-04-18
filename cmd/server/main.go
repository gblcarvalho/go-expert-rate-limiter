package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gblcarvalho/go-expert-rate-limiter/pkg/configs"
	"github.com/gblcarvalho/go-expert-rate-limiter/pkg/ratelimiter"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

func main() {
	config, err := configs.LoadConfig("../.", ".env")
	if err != nil {
		panic(err)
	}
	// store := ratelimiter.NewMemoryStore()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
	})
	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	store := ratelimiter.NewRedisStore(rdb)
	opts := ratelimiter.RateLimiterMiddlewareOpts{
		IPMaxRequests:    config.IPMaxRequests,
		TokenMaxRequests: config.TokenMaxRequests,
		TimeWindow:       time.Duration(config.TimeWindow * int64(time.Millisecond)),
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
