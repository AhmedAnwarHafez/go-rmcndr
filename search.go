package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetSearchHandler(c *fiber.Ctx) error {

	// query all users
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		return err
	}
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return err
	}

	// create a slice of users
	users := []User{}

	// iterate over each row
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Nickname, &user.IsPublic)
		if err != nil {
			return err
		}

		users = append(users, user)
	}

	fmt.Println(users)

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
