package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"go-rcmndr/routes"
)

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

	app.Get("/login", routes.Login)
	app.Get("/logout", routes.Logout)
	app.Get("/auth/github/callback", routes.GetAuthCallback)
	app.Get("/profile", routes.Profile)

	log.Fatal(app.Listen(":3000"))
}
