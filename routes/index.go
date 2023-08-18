package routes

import (
	"fmt"
	"go-rcmndr/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	session, err := store.Get(c)
	if err != nil {
		fmt.Println("error")
		return c.SendString("error")
	}

	userId, ok := session.Get("user_id").(int)
	if !ok {
		return c.Redirect("/login")
	}

	profile, err := db.GetProfileById(strconv.Itoa(userId))
	if err != nil {
		return c.SendString("error")
	}

	fmt.Println(profile)
	return c.SendString("Hello, World!")
}
