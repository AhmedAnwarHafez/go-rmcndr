package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var Connection = redis.NewClient(&redis.Options{
	Addr:     "192.168.1.184:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func GetProfileById(id string) (string, error) {
	key := "user:" + id
	return Connection.Get(ctx, key).Result()
}
