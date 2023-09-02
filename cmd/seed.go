package main

import (
	"fmt"
	"log"

	"context"
	"github.com/jaswdr/faker"
	"github.com/joho/godotenv"
)

type Song struct {
	ID      int64
	Title   string
	Artist  string
	Genre   string
	Link    string
	Comment string
}

type User struct {
	ID       int64
	Nickname string
	IsPublic bool
	Songs    []Song
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// seed := mathrand.NewSource(123)
	// f := faker.NewWithSeed(seed)

}

func followUser(ctx context.Context, userA string, userB string) error {
	return nil
}

func insertUsers(ctx context.Context, f faker.Faker) {
	for i := int64(0); i < 100; i++ {

		// create user
		user := User{
			ID:       int64(i),
			Nickname: f.Person().FirstNameFemale(),
			IsPublic: false,
		}

		key := fmt.Sprintf("users:%d", user.ID)
		fmt.Println(key)
	}
}
