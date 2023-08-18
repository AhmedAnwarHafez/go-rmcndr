package routes

import (
	"encoding/json"
	"go-rcmndr/db"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProfileData struct {
	UserId   int
	Nickname string
	Public   bool
}

func Profile(c *fiber.Ctx) error {
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
	val, err := db.GetProfileById(key)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	var profileData ProfileData
	err = json.Unmarshal([]byte(val), &profileData)
	if err != nil {
		log.Println("error unmarshalling profile data")
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Render("profile", fiber.Map{
		"Title":    "rcmndr - Profile",
		"Nickname": profileData.Nickname,
	}, "layouts/main")
}
