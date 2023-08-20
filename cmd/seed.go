package main

import (
	"fmt"
	"log"

	"context"
	"github.com/jaswdr/faker"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	mathrand "math/rand"
	"os"
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

func main() {

	var ctx = context.Background()
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

	insertUsers(rdb, ctx, f)
	// followUser(ctx, rdb, "1", "2")
}

func followUser(ctx context.Context, rdb *redis.Client, userA string, userB string) error {
	// Add userB to userA's following list
	if err := rdb.ZAdd(ctx, "following:"+userA, redis.Z{
		Score:  0,
		Member: "",
	}).Err(); err != nil {
		return err
	}

	// Add userA to userB's followers list
	if err := rdb.ZAdd(ctx, "followers:"+userB, redis.Z{Score: 0, Member: ""}).Err(); err != nil {
		return err
	}

	return nil
}

func insertRecommendations(rdb *redis.Client, ctx context.Context, f faker.Faker, userId int64) {

	randId := f.Int64Between(1, 10)

	for j := int64(0); j < randId; j++ {

		song := Song{
			Title:   f.Lorem().Word(),
			Artist:  f.Music().Author(),
			Genre:   f.Music().Genre(),
			Link:    f.Internet().URL(),
			Comment: f.Lorem().Sentence(10),
		}

		key := fmt.Sprintf("songs:%d", userId*11+j)
		err := rdb.HSet(ctx,
			key, "title", song.Title,
			"artist", song.Artist,
			"genre", song.Genre,
			"link", song.Link,
			"comment", song.Comment,
		).Err()

		if err != nil {
			log.Fatalf("Error setting song: %v", err)
		}

		err = rdb.ZAdd(ctx, fmt.Sprintf("recommendations:%d", userId), redis.Z{
			Score:  float64(j),
			Member: key,
		}).Err()

		if err != nil {
			log.Fatalf("Error setting recommendation: %v", err)
		}
	}
}

func insertUsers(rdb *redis.Client, ctx context.Context, f faker.Faker) {
	for i := int64(0); i < 100; i++ {

		// create user
		user := User{
			ID:       int64(i),
			Nickname: f.Person().FirstNameFemale(),
			IsPublic: false,
		}

		key := fmt.Sprintf("users:%d", user.ID)
		err := rdb.HSet(ctx, key, "nickname", user.Nickname, "public", user.IsPublic).Err()
		if err != nil {
			log.Fatalf("Error setting user: %v", err)
		}

		insertRecommendations(rdb, ctx, f, i)
	}
}
