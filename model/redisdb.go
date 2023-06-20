package model

import (
	"context"
	redis "github.com/go-redis/redis/v8"
)

var (
	Redisdb *redis.Client
	Ctx     context.Context
)

func InitRedisdb() {
	Ctx = context.Background()
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
