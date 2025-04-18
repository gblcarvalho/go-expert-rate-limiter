package ratelimiter

import (
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

func (rl *RateLimiter) Allow(key string, maxRequests int) bool {
	count := rl.store.IncrementOrReset(key, rl.TimeWindow)
	return count <= maxRequests
}
