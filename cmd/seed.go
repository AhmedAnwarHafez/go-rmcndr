package main

import (
	"database/sql"
	"github.com/jaswdr/faker"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"log"
	rand "math/rand"
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
	// create a new table called recommendations
	// with the following columns:
	// id, title, artist, genre, url
	// and insert some data
	// then select all data from the table
	// and print it to the console
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	// create table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS recommendations (id INTEGER PRIMARY KEY, title TEXT, artist TEXT, genre TEXT, url TEXT, comment TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// delete all recommendations
	_, err = db.Exec("DELETE FROM recommendations")
	if err != nil {
		log.Fatal(err)
	}

	// create a new faker instance
	fake := faker.New()

	for i := 0; i < 10; i++ {

		// create a new recommendation
		recommendation := Song{
			Title:   fake.Music().Name(),
			Artist:  fake.Person().Name(),
			Genre:   fake.Music().Genre(),
			Link:    fake.Internet().URL(),
			Comment: fake.Lorem().Sentence(rand.Intn(30) + 10),
		}

		// insert the recommendation into the database
		_, err = db.Exec("INSERT INTO recommendations (title, artist, genre, url, comment) VALUES (?, ?, ?, ?, ?)", recommendation.Title, recommendation.Artist, recommendation.Genre, recommendation.Link, recommendation.Comment)
	}

	if err != nil {
		log.Fatal(err)
	}
}
