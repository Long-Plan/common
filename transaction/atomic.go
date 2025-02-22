package transaction

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ErrResourceLocked = errors.New("resource is already locked")

type atomic struct {
	Redis *redis.Client
	Key   string
	TTL   time.Duration
}

// NewAtomic creates a new atomic  handler.
func NewAtomic(redis *redis.Client, key string, ttl time.Duration) IAtomic {
	return &atomic{
		Redis: redis,
		Key:   key,
		TTL:   ttl,
	}
}

type IAtomic interface {
	Lock(ctx context.Context, id string) error
	Release(ctx context.Context, id string) error
}

func (a *atomic) Lock(ctx context.Context, id string) error {
	success, err := a.Redis.SetNX(ctx, fmt.Sprintf("%s:%s:%s", "ATOMIC", a.Key, id), "LOCKED", a.TTL).Result()
	if err != nil {
		return fmt.Errorf("failed to set lock in Redis: %w", err)
	}

	if !success {
		return ErrResourceLocked
	}

	return nil
}

func (a *atomic) Release(ctx context.Context, id string) error {
	if err := a.Redis.Del(ctx, fmt.Sprintf("%s:%s:%s", "ATOMIC", a.Key, id)).Err(); err != nil {
		return fmt.Errorf("failed to release lock in Redis: %w", err)
	}

	return nil
}
