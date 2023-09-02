package main

import (
	"fmt"

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

	fmt.Println(userId)
	return c.SendString("Hello, World!")
}
