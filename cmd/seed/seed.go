package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jaswdr/faker"
	"github.com/joho/godotenv"
	mathrand "math/rand"
	"os"

	"github.com/redis/go-redis/v9"
)

type Song struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Artist  string `json:"artist"`
	Genre   string `json:"genre"`
	Link    string `json:"link"`
	Comment string `json:"comment"`
}

type User struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	IsPublic bool   `json:"is_public"`
	Songs    []Song `json:"songs"`
}

var ctx = context.Background()

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	seed := mathrand.NewSource(123)
	f := faker.NewWithSeed(seed)

	for i := 0; i < 100; i++ {

		song := Song{
			ID:      int64(i),
			Title:   f.Lorem().Word(),
			Artist:  f.Music().Author(),
			Genre:   f.Music().Genre(),
			Link:    f.Internet().URL(),
			Comment: f.Lorem().Sentence(10),
		}

		// create user
		user := User{
			ID:       int64(i),
			Nickname: f.Person().FirstNameFemale(),
			IsPublic: false,
			Songs:    []Song{song},
		}

		fmt.Printf("%+v\n", song)

		j, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}

		err = rdb.Set(ctx, fmt.Sprintf("user:%d", user.ID), j, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}
