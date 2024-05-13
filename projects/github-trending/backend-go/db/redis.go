package db

import (
	"context"
	"fmt"
	"github-trending-api/logger"
	"log"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	RedisClient *redis.Client
	Host        string
	Port        int
	Password    string
}

func (r *Redis) Connect() {
	r.RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       0,          // use default DB
	})

	if err := r.RedisClient.Ping(context.Background()).Err(); err != nil {
		logger.Error(err.Error())
	}
	log.Println("Connected to REDIS")
}

func (r *Redis) Close() {
	log.Println("Disconnected to REDIS")
	r.RedisClient.Close()
}
