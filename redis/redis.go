package redis

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Address    string
	Password   string
	Database   int    // optional
	AuthString string // optional
}

func NewRedisClient(ctx context.Context, config *Config) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         config.Address,
		Password:     config.Password,
		DB:           config.Database,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	})
	if config.AuthString != "" {
		if _, err := redisClient.Do(ctx, "AUTH", config.AuthString).Result(); err != nil {
			redisClient.Close()
			return nil, err
		}
	}
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, errors.New("Cannot connect to redis: " + err.Error())
	}
	return redisClient, nil
}
