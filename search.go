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
	rows, err := db.Query(`SELECT u.nickname as Nickname, u.is_public as IsPublic, g.name as GenreName, COUNT(r.id) AS GenereCount
FROM recommendations r
JOIN users u ON r.user_id = u.id
JOIN genres g ON r.genre_id = g.id
WHERE u.nickname LIKE ?
GROUP BY g.id`, "%br%")

	if err != nil {
		return err
	}

	defer rows.Close()

	type Genre struct {
		GenreName string
		Count     int64
	}

	type User struct {
		Id       int64
		Nickname string
		IsPublic bool
		Genres   []Genre
	}

	usersMap := make(map[string]*User)

	for rows.Next() {
		var user User
		var genre Genre
		err := rows.Scan(&user.Nickname, &user.IsPublic, &genre.GenreName, &genre.Count)
		if err != nil {
			return err
		}
		// Check if user already exists in the map
		if _, exists := usersMap[user.Nickname]; !exists {
			usersMap[user.Nickname] = &User{Nickname: user.Nickname}
		}

		// Append the genre to the user's genres
		usersMap[user.Nickname].Genres = append(usersMap[user.Nickname].Genres, genre)
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
		"Users": usersMap,
	}, "layouts/main")
}
