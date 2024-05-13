package helper

import (
	"context"
	"github-trending-api/db"
	"github-trending-api/logger"
	"time"
)

// Set data to redis cache.
func SetCache(ctx context.Context, key string, value string) error {
	// Connect to redis
	redis := &db.Redis{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "",
	}
	redis.Connect()
	defer redis.Close()

	err := redis.RedisClient.Set(ctx, key, value, time.Minute*1).Err()
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// Get data from redis cache
func GetCache(ctx context.Context, key string) (string, error) {
	// Connect to redis
	redis := &db.Redis{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "",
	}
	redis.Connect()
	defer redis.Close()

	val, err := redis.RedisClient.Get(ctx, key).Result()
	if err != nil {
		logger.Error(err)
		return "", err
	}

	return val, nil
}
