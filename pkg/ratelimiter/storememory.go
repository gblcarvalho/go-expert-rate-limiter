package ratelimiter

import (
	"sync"
	"time"
)

type MemoryStoreRate struct {
	count  int
	initAt time.Time
}
type MemoryStore struct {
	rates map[string]MemoryStoreRate
	mux   sync.Mutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		rates: make(map[string]MemoryStoreRate),
	}
}

func (s *MemoryStore) IncrementOrReset(key string, duration time.Duration) int {
	s.mux.Lock()
	defer s.mux.Unlock()

	rate, ok := s.rates[key]
	now := time.Now()

	if !ok || now.Sub(rate.initAt) > duration {
		rate = MemoryStoreRate{
			count:  0,
			initAt: now,
		}
	}

	rate.count++
	s.rates[key] = rate
	return rate.count
}
