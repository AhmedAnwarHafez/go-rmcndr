package main

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func GetFriendDetails(c *fiber.Ctx) error {

	// decode params to string
	nickname := c.Params("nickname")
	nickname, err := url.QueryUnescape(nickname)
	fmt.Println(nickname)

	// get user from db
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	var bio sql.NullString
	err = db.QueryRow("SELECT bio FROM users WHERE nickname = ?", nickname).Scan(&bio)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Treat NULL as an empty string
	if !bio.Valid {
		bio.String = ""
	}

	var recommendations []Song
	rows, err := db.Query("SELECT id, user_id, title, artist, genre_id, url, comment FROM recommendations WHERE user_id = (SELECT id FROM users WHERE nickname = ?)", nickname)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	for rows.Next() {
		var r Song
		err = rows.Scan(&r.Id, &r.UserId, &r.Title, &r.Artist, &r.GenreId, &r.Link, &r.Comment)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		recommendations = append(recommendations, r)
	}

	return c.Render("friend", fiber.Map{
		"Title":           nickname,
		"Nickname":        nickname,
		"Bio":             bio.String,
		"Recommendations": recommendations,
	}, "layouts/main")
}
