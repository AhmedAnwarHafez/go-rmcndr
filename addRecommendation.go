package main

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func GetAddRecommendation(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return c.Render("add-recommendation", fiber.Map{
		"Title": "Home - rcmndr",
	}, "layouts/main")
}

func PostAddRecommendation(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// get form values
	title := c.FormValue("title")
	artist := c.FormValue("artist")
	genre := c.FormValue("genre")
	url := c.FormValue("url")
	comment := c.FormValue("comment")

	// insert
	stmt, err := db.Prepare("INSERT INTO recommendations(title, artist, genre, url, comment) values(?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(title, artist, genre, url, comment)
	if err != nil {
		log.Fatal(err)
	}

	return c.Redirect("/add-recommendation")
}
