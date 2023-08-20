package main

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var Connection = redis.NewClient(&redis.Options{
	Addr:     "192.168.1.184:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func GetProfileById(id string) (User, error) {
	key := "user:" + id

	val, err := Connection.Get(ctx, key).Result()
	if err != nil {
		return User{}, err
	}

	var u User
	err = json.Unmarshal([]byte(val), &u)

	return u, err
}

func GetSongs(id string) ([]Song, error) {
	profile, err := GetProfileById(id)

	return profile.Songs, err
}
