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

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/public", "./public")

	app.Get("/login", LoginHandler)
	app.Get("/logout", LogoutHandler)
	app.Get("/auth/github/callback", GetAuthCallbackHanlder)
	app.Get("/profile", GetProfileHandler)
	app.Get("/add-recommendation", GetAddRecommendation)
	app.Post("/add-recommendation", PostAddRecommendation)
	app.Get("/search", GetSearchHandler)
	app.Get("/", GetSongsHandler)

	log.Fatal(app.Listen(":3000"))
}
