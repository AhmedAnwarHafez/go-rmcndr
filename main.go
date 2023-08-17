package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"go-rcmndr/routes/auth"
	"golang.org/x/oauth2"
)

// session store
var store = session.New()

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/public", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Title": "rcmndr",
		}, "layouts/main")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		config := auth.GetConfig()
		authURL := config.AuthCodeURL("", oauth2.AccessTypeOffline)
		return c.Redirect(authURL)
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer sess.Save()

		sess.Destroy()
		return c.Redirect("/")
	})

	app.Get("/auth/github/callback", func(c *fiber.Ctx) error {
		code := c.Query("code")
		config := auth.GetConfig()
		token, err := config.Exchange(c.Context(), code)

		user, err := auth.GetUserInfo(token, c, config)
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
		return c.Redirect("/protected")
	})
	// protected rout
	app.Get("/protected", func(c *fiber.Ctx) error {
		// read user_id from session
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		userID := sess.Get("user_id")
		if userID == nil {
			// user not logged in
			return c.Redirect("/auth/github/login")
		}

		return c.Render("profile", fiber.Map{
			"Title":  "rcmndr - Profile",
			"UserId": userID,
		}, "layouts/main")
	})

	log.Fatal(app.Listen(":3000"))
}
