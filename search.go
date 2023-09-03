package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func GetSearchHandler(c *fiber.Ctx) error {

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		return err
	}
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return err
	}

	users := []User{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Nickname, &user.IsPublic)
		if err != nil {
			return err
		}

		users = append(users, user)
	}

	q := c.Query("q", "default")

	// means that htmx triggered this request
	if c.Get("Hx-Trigger-Name") == "q" {

		return c.Render("search-result", fiber.Map{
			"Title": "rcmndr",
			"Query": q,
		}, "layouts/main")
	}

	return c.Render("search", fiber.Map{
		"Title": "rcmndr",
		"Query": q,
		"Users": users,
	}, "layouts/main")
}
