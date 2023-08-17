package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"go-rcmndr/routes/utils"
	"golang.org/x/oauth2"
)

// session store
var store = session.New()

func GetAuthCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	config := utils.GetConfig()
	token, err := config.Exchange(c.Context(), code)

	user, err := utils.GetUserInfo(token, c, config)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer sess.Save()

	log.Println(user.Login)

	sess.Set("user_id", user.Login)
	return c.Redirect("/profile")
}

func Logout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer sess.Save()

	sess.Destroy()
	return c.Redirect("/")
}

func Login(c *fiber.Ctx) error {
	config := utils.GetConfig()
	authURL := config.AuthCodeURL("", oauth2.AccessTypeOffline)
	return c.Redirect(authURL)
}
