package main

import (
	"bytes"
	"html/template"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new engine
	engine := html.New("./views", ".html")

	engine.AddFunc("textbox", func(class string, placeholder template.HTML) template.HTML {
		var buf bytes.Buffer
		engine.Templates.ExecuteTemplate(&buf, "textbox", map[string]interface{}{
			"Class":       class,
			"Placeholder": placeholder,
		})
		log.Println(buf.String())
		return template.HTML(buf.String())
	})

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/public", "./public")

	app.Get("/", func(c *fiber.Ctx) error {

		return Index(c)
		// return c.Render("index", fiber.Map{
		// 	"Title": "rcmndr",
		// }, "layouts/main")
	})

	app.Get("/login", LoginHandler)
	app.Get("/logout", LogoutHandler)
	app.Get("/auth/github/callback", GetAuthCallbackHanlder)
	app.Get("/profile", ProfileHandler)
	app.Get("/", HomeHandler)

	app.Get("/search", func(c *fiber.Ctx) error {

		q := c.Query("q", "default")

		// means that htmx triggered this request
		if c.Get("Hx-Trigger-Name") == "q" {

			return c.Render("search-result", fiber.Map{
				"Title": "rcmndr",
				"Query": q,
			}, "layouts/main")
		}

		return c.Render("search", fiber.Map{
			"Title": "rcmndr",
			"Query": q,
		}, "layouts/main")
	})

	log.Fatal(app.Listen(":3000"))
}
