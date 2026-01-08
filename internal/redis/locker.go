package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Locker struct {
	client *redis.Client
}

func NewLocker(addr string) *Locker {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &Locker{client: rdb}
}

func (l *Locker) AcquireLock(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return l.client.SetNX(ctx, "lock:"+key, "1", expiration).Result()
}

func (l *Locker) ReleaseLock(ctx context.Context, key string) error {
	return l.client.Del(ctx, "lock:"+key).Err()
}
