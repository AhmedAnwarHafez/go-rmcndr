package main

import (
	"fmt"
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

	fmt.Println(userId)
	profile, err := GetProfileById(strconv.Itoa(userId))
	if err != nil {
		fmt.Println(err)
		return c.SendString("unable to get profile")
	}

	fmt.Println(profile)
	return c.SendString("Hello, World!")
}
