package db

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func GetProfileById(key string) (string, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.20.244:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb.Get(ctx, key).Result()
}
