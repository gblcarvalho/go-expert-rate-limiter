package ratelimiter

import (
	"log"
	"time"
)

type RateLimiter struct {
	TimeWindow time.Duration
	store      RateLimiterStoreInterface
}

func NewRateLimiter(store RateLimiterStoreInterface, TimeWindow time.Duration) (*RateLimiter, error) {
	return &RateLimiter{
		TimeWindow: TimeWindow,
		store:      store,
	}, nil
}

func (rl *RateLimiter) Allow(key string, maxRequests int64) bool {
	count, err := rl.store.IncrementOrReset(key, rl.TimeWindow)
	if err != nil {
		log.Println(err)
		return false
	}
	return count <= maxRequests
}
