package main

import (
	"database/sql"
	"log"
	rand "math/rand"

	"github.com/jaswdr/faker"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type Recommendation struct {
	ID      int64
	UserId  int64
	GenreId int64
	Title   string
	Artist  string
	Link    string
	Comment string
}

type User struct {
	Id       int64
	Nickname string
	Bio      string
	IsPublic bool
	Songs    []Recommendation
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
		return
	}

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
		return
	}

	// create users table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, nickname TEXT, bio TEXT, is_public BOOLEAN)")
	if err != nil {
		log.Fatal(err)
		return
	}

	// create genres table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS genres (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatal(err)
		return
	}

	// create table for recommendations with foreign key to genres and users
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS recommendations (id INTEGER PRIMARY KEY, title TEXT, artist TEXT, genre_id INTEGER, url TEXT, comment TEXT, user_id INTEGER, FOREIGN KEY(genre_id) REFERENCES genres(id), FOREIGN KEY(user_id) REFERENCES users(id))")
	if err != nil {
		log.Fatal(err)
		return
	}

	// delete all recommendations
	_, err = db.Exec("DELETE FROM recommendations")
	if err != nil {
		log.Fatal(err)
		return
	}

	// delete all users
	_, err = db.Exec("DELETE FROM users")
	if err != nil {
		log.Fatal(err)
		return
	}

	// delete all genres
	_, err = db.Exec("DELETE FROM genres")
	if err != nil {
		log.Fatal(err)
		return
	}

	// create a new faker instance with seed 123
	fake := faker.NewWithSeed(rand.NewSource(123))

	var userIds []int64
	// insert 10 users into the database
	for i := 0; i < 100; i++ {
		// create a new user
		u := User{
			Nickname: fake.Person().Name(),
			Bio:      fake.Lorem().Sentence(10),
			IsPublic: true,
		}

		// insert the user into the database
		sql, err := db.Exec("INSERT INTO users (nickname, is_public) VALUES (?, ?)", u.Nickname, u.IsPublic)
		if err != nil {
			log.Fatal(err)
			return
		}

		// get the id of the user
		userId, err := sql.LastInsertId()
		if err != nil {
			log.Fatal(err)
			return
		}
		userIds = append(userIds, userId)
	}

	var genreIds []int64
	for _, genre := range GetGenres() {
		// create a new genre

		// insert the genre into the database
		sql, err := db.Exec("INSERT INTO genres (name) VALUES (?)", genre)
		if err != nil {
			log.Fatal(err)
			return
		}

		id, err := sql.LastInsertId()
		if err != nil {
			log.Fatal(err)
			return
		}
		genreIds = append(genreIds, id)
	}

	for i := 0; i < 500; i++ {
		// create a new r
		r := Recommendation{
			UserId:  userIds[rand.Intn(len(userIds))],
			GenreId: genreIds[rand.Intn(len(genreIds))],
			Title:   fake.Music().Name(),
			Artist:  fake.Person().Name(),
			Link:    fake.Internet().URL(),
			Comment: fake.Lorem().Sentence(rand.Intn(30) + 10),
		}

		// insert the recommendation into the database
		_, err = db.Exec("INSERT INTO recommendations (user_id, genre_id, title, artist, url, comment) VALUES (?, ?, ?, ?, ?, ?)", r.UserId, r.GenreId, r.Title, r.Artist, r.Link, r.Comment)

		if err != nil {
			log.Fatal(err)
			return
		}
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	var me = User{
		Id:       7552088,
		Nickname: "me",
		Bio:      fake.Lorem().Sentence(10),
		IsPublic: true,
	}

	// insert user into the database
	_, err = db.Exec("INSERT INTO users (id, nickname, bio, is_public) VALUES (?, ?, ?, ?)", me.Id, me.Nickname, me.Bio, me.IsPublic)
	if err != nil {
		log.Fatal(err)
		return
	}

	for i := 0; i < 20; i++ {
		// create a new r
		r := Recommendation{
			UserId:  me.Id,
			GenreId: genreIds[rand.Intn(len(genreIds))],
			Title:   fake.Music().Name(),
			Artist:  fake.Person().Name(),
			Link:    fake.Internet().URL(),
			Comment: fake.Lorem().Sentence(rand.Intn(30) + 10),
		}

		// insert the recommendation into the database
		_, err = db.Exec("INSERT INTO recommendations (user_id, genre_id, title, artist, url, comment) VALUES (?, ?, ?, ?, ?, ?)", r.UserId, r.GenreId, r.Title, r.Artist, r.Link, r.Comment)

		if err != nil {
			log.Fatal(err)
			return
		}
	}

}

func GetGenres() []string {
	musicGenres := []string{
		"Dance-pop", "Dance-punk", "Dance-rock", "Electro",
		"Power pop", "Powerviolence", "Progressive breaks", "Progressive drum & bass", "Progressive folk", "Progressive house", "Progressive metal", "Progressive rock", "Progressive techno",
		"Progressive", "Psybreaks", "Psychedelic folk", "Psychedelic rock", "Psychedelic trance", "Psychobilly", "Psyprog", "Punk jazz", "Punk rock",
		"Raga rock", "Ragga-jungle", "Raggacore", "Ragtime", "Rap metal", "Rap rock", "Rapcore", "Riot grrrl", "Rock and roll",
		"Sufi rock", "Sung poetry", "Suomisaundi", "Surf rock", "Swing house", "Swing", "Symphonic metal", "Synthcore", "Synthpop",
		"Tecno brega", "Terrorcore", "Third stream", "Thrash metal", "Thrashcore", "Toytown Techno", "Trad jazz", "Traditional doom", "Trance",
	}

	return musicGenres
}
