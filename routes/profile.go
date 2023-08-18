package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go-rcmndr/db"
	"log"
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

	val, err := db.GetProfileById("7552088")
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
		"UserId":   userId,
		"Nickname": profileData.Nickname,
	}, "layouts/main")
}
