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
