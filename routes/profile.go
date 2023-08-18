package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Profile(c *fiber.Ctx) error {
	// read user_id from session
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	userID := sess.Get("user_id")
	if userID == nil {
		// user not logged in
		return c.Redirect("/login")
	}

	return c.Render("profile", fiber.Map{
		"Title":  "rcmndr - Profile",
		"UserId": userID,
	}, "layouts/main")
}
