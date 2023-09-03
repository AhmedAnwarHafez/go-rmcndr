package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func GetSongsHandler(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	sessionUser := sess.Get("user_id")
	if sessionUser == nil {
		return c.Redirect("/login")
	}

	fmt.Println(sessionUser)
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// select all
	rows, err := db.Query("SELECT id, title, artist, genre, url, comment FROM recommendations WHERE user_id = ?", sessionUser)
	if err != nil {
		log.Fatal(err)
	}

	var recommendations []Song

	if rows.Next() == false {
		log.Println("No recommendations found")
		return c.Render("list-recommendations", fiber.Map{
			"Title":           "Nothing - rcmndr",
			"Recommendations": recommendations,
		}, "layouts/main")
	}

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

		recommendations = append(recommendations, Song{
			Id:      id,
			Title:   title,
			Artist:  artist,
			Genre:   genre,
			Link:    url,
			Comment: comment,
		})
	}

	return c.Render("list-recommendations", fiber.Map{
		"Title":           "Home - rcmndr",
		"Recommendations": recommendations,
	}, "layouts/main")
}
