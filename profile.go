package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func GetProfileHandler(c *fiber.Ctx) error {
	// read user_id from session
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	userId, ok := sess.Get("user_id").(int)
	if !ok {
		// user not logged in
		return c.SendString("user not logged in")
	}

	// get user from db
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	var user User
	err = db.QueryRow("SELECT id, nickname, bio, is_public FROM users WHERE id = ?", userId).Scan(&user.Id, &user.Nickname, &user.Bio, &user.IsPublic)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Render("profile", fiber.Map{
		"Title":    "rcmndr - Profile",
		"Nickname": user.Nickname,
		"Bio":      user.Bio,
		"IsPublic": user.IsPublic,
	}, "layouts/main")
}
