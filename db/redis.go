package main

import (
	"context"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_URL"),
	Password: "", // no password set
	DB:       0,  // use default DB
})

func GetProfileById(id int) (string, error) {
	// convert int to string
	key := strconv.Itoa(id)
	return rdb.Get(ctx, key).Result()
}
