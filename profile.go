package main

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProfileData struct {
	UserId   int
	Nickname string
	Public   bool
}

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

	key := strconv.Itoa(userId)
	val, err := GetProfileById(key)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	var profileData ProfileData
	err = json.Unmarshal([]byte(val), &profileData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Render("profile", fiber.Map{
		"Title":    "rcmndr - Profile",
		"Nickname": profileData.Nickname,
	}, "layouts/main")
}