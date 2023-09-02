package main

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func HomeHandler(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// select all
	rows, err := db.Query("SELECT * FROM recommendations")
	if err != nil {
		log.Fatal(err)
	}

	var recommendations []Song

	// iterate over rows
	for rows.Next() {
		var id int64
		var title string
		var artist string
		var genre string
		var url string
		var comment string
		err = rows.Scan(&id, &title, &artist, &genre, &url, &comment)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, title, artist, genre, url, comment)

		recommendations = append(recommendations, Song{
			Id:      id,
			Title:   title,
			Artist:  artist,
			Genre:   genre,
			Link:    url,
			Comment: comment,
		})
	}

	return c.Render("recommendations", fiber.Map{
		"Title":           "Home - rcmndr",
		"Recommendations": recommendations,
	}, "layouts/main")
}
