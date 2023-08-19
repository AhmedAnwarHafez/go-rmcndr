package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"golang.org/x/oauth2"
)

// Define Redis connection configuration
var redisConfig = redis.Config{
	Host:     "192.168.1.184", // Redis server IP and
	Port:     6379,            // Redis server port
	Username: "",              // Redis username (if applicable)
	Password: "",              // Redis password (if applicable)
	Database: 0,               // Redis database number
}

// session store
var store = session.New(
	session.Config{
		Storage: redis.New(redisConfig),
	})

func GetAuthCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	//
	config := GetConfig()
	token, err := config.Exchange(c.Context(), code)

	user, err := GetUserInfo(token, c, config)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	sess.Set("user_id", user.ID)
	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Redirect("/profile")
}

func Logout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if err := sess.Destroy(); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Redirect("/")
}

func Login(c *fiber.Ctx) error {
	config := GetConfig()
	authURL := config.AuthCodeURL("", oauth2.AccessTypeOffline)
	return c.Redirect(authURL)
}
