package main

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"go-rcmndr/routes"
	"html/template"
	"log"
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

		return routes.Index(c)
		// return c.Render("index", fiber.Map{
		// 	"Title": "rcmndr",
		// }, "layouts/main")
	})

	app.Get("/login", routes.Login)
	app.Get("/logout", routes.Logout)
	app.Get("/auth/github/callback", routes.GetAuthCallback)
	app.Get("/profile", routes.Profile)
	app.Get("/profile", routes.Profile)

	log.Fatal(app.Listen(":3000"))
}
