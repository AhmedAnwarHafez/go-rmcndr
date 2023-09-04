package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

// Middleware to check if the user is authenticated
func Protect(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil || sess.Get("user_id") == nil {
		return c.Redirect("/login")
	}

	log.Println("User is authenticated")
	return c.Next()
}

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
	app.Get("/profile", Protect, GetProfileHandler)
	app.Get("/add-recommendation", Protect, GetAddRecommendation)
	app.Post("/add-recommendation", Protect, PostAddRecommendation)
	app.Get("/friend/:nickname", GetFriendDetails)
	app.Get("/search", Protect, GetSearchHandler)
	app.Get("/", Protect, GetSongsHandler)

	log.Fatal(app.Listen(":3000"))
}
