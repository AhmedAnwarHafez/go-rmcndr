package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jaswdr/faker"
	"go-rcmndr/db"
	mathrand "math/rand"
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
}

var ctx = context.Background()

func main() {
	// connect to redis
	rdb := db.Connection
	seed := mathrand.NewSource(123)
	f := faker.NewWithSeed(seed)

	for i := 0; i < 100; i++ {
		// create user
		user := User{
			ID:       int64(i),
			Nickname: f.Person().FirstNameFemale(),
			IsPublic: false,
		}

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
