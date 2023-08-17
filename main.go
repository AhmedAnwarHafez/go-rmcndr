package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/template/html/v2"
)

func main() {

	// Create a new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/public", "./public")
	cfg := basicauth.Config{
		Users: map[string]string{
			"banana": "mango",
		},
	}

	app.Use(basicauth.New(cfg))
	// Group protected routes that require Basic Authentication
	protected := app.Group("", basicauth.New(cfg))
	protected.Get("/protected", func(c *fiber.Ctx) error {
		return c.SendString("Protected route, authentication required!")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Title": "rcmndr",
		}, "layouts/main")
	})

	log.Fatal(app.Listen(":3000"))
}
