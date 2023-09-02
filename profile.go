package main

import (
	"github.com/gofiber/fiber/v2"
)

func ProfileHandler(c *fiber.Ctx) error {
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

	return c.Render("profile", fiber.Map{
		"Title":    "rcmndr - Profile",
		"Nickname": userId,
	}, "layouts/main")
}
