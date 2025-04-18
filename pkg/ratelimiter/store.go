package ratelimiter

import "time"

type RateLimiterStoreInterface interface {
	IncrementOrReset(key string, duration time.Duration) (int64, error)
}
