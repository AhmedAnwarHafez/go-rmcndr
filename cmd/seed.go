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
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, nickname TEXT, is_public BOOLEAN)")
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

	for i := 0; i < 300; i++ {
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
		IsPublic: true,
	}

	// insert user into the database
	_, err = db.Exec("INSERT INTO users (id, nickname, is_public) VALUES (?, ?, ?)", me.Id, me.Nickname, me.IsPublic)
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
		"2-step", "4-beat", "Acid breaks", "Acid house", "Acid jazz", "Acid rock", "Acid techno", "Acid trance", "Aggrotech",
		"Alternative dance", "Alternative metal", "Alternative rock", "Ambient dub", "Ambient house", "Ambient techno", "Ambient", "Anarcho punk", "Anti-folk",
		"Art punk", "Art rock", "Asian Underground", "Avant-garde jazz", "Baggy", "Balearic Beat", "Baltimore Club", "Bassline", "Beat music",
		"Bebop", "Big beat", "Bitpop", "Black metal", "Boogie-woogie", "Boogie", "Bossa nova", "Bouncy house", "Bouncy techno",
		"Breakbeat hardcore", "Breakbeat", "Breakcore", "Breakstep", "British dance", "Britpop", "Broken beat", "Bubblegum dance", "Canterbury scene",
		"Cape jazz", "Celtic metal", "Celtic punk", "Celtic", "Chamber jazz", "Chicago house", "Chill out", "Chillwave", "Chinese rock",
		"Chiptune", "Christian metal", "Christian punk", "Christian rock", "Classic trance", "Coldwave", "Contemporary folk", "Continental Jazz", "Cool jazz",
		"Cosmic disco", "Cowpunk", "Crossover jazz", "Crossover thrash", "Crunk", "Crust punk", "Crustgrind", "Cybergrind", "D-beat",
		"Dance-pop", "Dance-punk", "Dance-rock", "Dark ambient", "Dark cabaret", "Dark electro", "Dark psytrance", "Dark Wave", "Darkcore",
		"Darkside jungle", "Darkstep", "Death industrial", "Death metal", "Deathcore", "Deathrock", "Deep house", "Desert rock", "Detroit techno",
		"Digital hardcore", "Disco house", "Disco polo", "Disco", "Diva house", "Dixieland", "Djent", "Doom metal", "Doomcore",
		"Downtempo", "Dream house", "Dream pop", "Dream trance", "Drone metal", "Drone", "Drum and bass", "Drumfunk", "Drumstep",
		"Dub", "Dubstep", "Dubstyle", "Dubtronica", "Dunedin Sound", "Dutch house", "EDM", "Electro backbeat", "Electro house",
		"Electro-grime", "Electro-industrial", "Electro", "Electroacoustic", "Electroclash", "Electronic art music", "Electronic rock", "Electronica", "Electronicore",
		"Electropop", "Electropunk", "Emo", "Epic doom", "Ethereal wave", "Ethnic electronica", "Euro disco", "Eurobeat", "Eurodance",
		"European free jazz", "Europop", "Experimental rock", "Filk", "Florida breaks", "Folk metal", "Folk punk", "Folk rock", "Folk",
		"Folktronica", "Freak folk", "Freakbeat", "Free tekno", "Freestyle house", "Freestyle", "French house", "Full on", "Funeral doom",
		"Funk metal", "Funky house", "Funky", "Futurepop", "Gabber", "Garage punk", "Garage rock", "Ghetto house", "Ghettotech",
		"Glam metal", "Glam rock", "Glitch", "Goregrind", "Gothic metal", "Gothic rock", "Grime", "Grindcore", "Groove metal",
		"Grunge", "Happy hardcore", "Hard bop", "Hard NRG", "Hard rock", "Hard trance", "Hardbag", "Hardcore punk", "Hardcore/Hard dance",
		"Hardstep", "Hardstyle", "Heavy metal", "Hi-NRG", "Hip house", "Horror punk", "House", "IDM", "Illbient",
		"Indie folk", "Indie pop", "Indie rock", "Indietronica", "Industrial folk", "Industrial metal", "Industrial rock", "Industrial", "Intelligent drum and bass",
		"Italo dance", "Italo disco", "Italo house", "Japanoise", "Jazz blues", "Jazz fusion", "Jazz rap", "Jazz rock", "Jazz-funk",
		"Jump-Up", "Jumpstyle", "Krautrock", "Laptronica", "Latin house", "Latin jazz", "Liquid funk", "Livetronica", "Lowercase",
		"Lo-fi", "Madchester", "Mainstream jazz", "Makina", "Math rock", "Mathcore", "Medieval metal", "Melodic death metal", "Metalcore",
		"Minimal house/Microhouse", "Minimal", "Modal jazz", "Moombahton", "Neo-bop jazz", "Neo-psychedelia", "Neo-swing", "Neofolk", "Neurofunk",
		"New Beat", "New jack swing", "New prog", "New rave", "New wave", "New-age", "Nintendocore", "No wave", "Noise pop",
		"Noise rock", "Noise", "Noisegrind", "Nortec", "Novelty ragtime", "Nu jazz", "Nu metal", "Nu skool breaks", "Nu-disco",
		"Oldschool jungle", "Orchestral jazz", "Orchestral Uplifting", "Paisley Underground", "Pop punk", "Pop rock", "Post-bop", "Post-Britpop", "Post-disco",
		"Post-grunge", "Post-hardcore", "Post-metal", "Post-punk revival", "Post-punk", "Post-rock", "Power electronics", "Power metal", "Power noise",
		"Power pop", "Powerviolence", "Progressive breaks", "Progressive drum & bass", "Progressive folk", "Progressive house", "Progressive metal", "Progressive rock", "Progressive techno",
		"Progressive", "Psybreaks", "Psychedelic folk", "Psychedelic rock", "Psychedelic trance", "Psychobilly", "Psyprog", "Punk jazz", "Punk rock",
		"Raga rock", "Ragga-jungle", "Raggacore", "Ragtime", "Rap metal", "Rap rock", "Rapcore", "Riot grrrl", "Rock and roll",
		"Rock in Opposition", "Sadcore", "Sambass", "Screamo", "Shibuya-kei", "Shoegaze", "Ska jazz", "Ska punk", "Skate punk",
		"Skweee", "Slowcore", "Sludge metal", "Smooth jazz", "Soft rock", "Soul jazz", "Sound art", "Southern rock", "Space disco",
		"Space house", "Space rock", "Speed garage", "Speed metal", "Speedcore", "Stoner rock", "Straight-ahead jazz", "Street punk", "Stride jazz",
		"Sufi rock", "Sung poetry", "Suomisaundi", "Surf rock", "Swing house", "Swing", "Symphonic metal", "Synthcore", "Synthpop",
		"Synthpunk", "Tech house", "Tech trance", "Technical death metal", "Techno-DNB", "Techno-folk", "Techno", "Technopop", "Techstep",
		"Tecno brega", "Terrorcore", "Third stream", "Thrash metal", "Thrashcore", "Toytown Techno", "Trad jazz", "Traditional doom", "Trance",
		"Trap", "Tribal house", "Trip hop", "Turbofolk", "Twee Pop", "Uplifting trance", "Vaporwave", "Viking metal", "Vocal house",
		"Vocal jazz", "Vocal trance", "West Coast jazz", "Western", "Witch House/Drag", "World fusion", "Worldbeat", "Yacht rock", "Yorkshire Bleeps and Bass",
	}

	return musicGenres
}
