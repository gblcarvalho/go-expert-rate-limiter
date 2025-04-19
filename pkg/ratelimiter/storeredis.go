package ratelimiter

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(rdb *redis.Client) *RedisStore {
	return &RedisStore{
		client: rdb,
	}
}

func (s *RedisStore) IncrementOrReset(key string, duration time.Duration) (int64, error) {
	ctx := context.Background()
	count, err := s.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	if count == 1 {
		s.client.Expire(ctx, key, duration)
	}
	return count, nil
}
