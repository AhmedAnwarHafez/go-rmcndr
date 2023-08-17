package main

import (
	"encoding/json"
	"log"
	// "net/http"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type GitHubUser struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
}

func getUserInfo(token *oauth2.Token, ctx *fiber.Ctx, config oauth2.Config) (*GitHubUser, error) {
	client := config.Client(ctx.Context(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	user := &GitHubUser{}
	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func main() {

	// Create a new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/public", "./public")

	oauth2Config := oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:3000/auth/github/callback",
		Scopes:       []string{}, // Add required scopes here
		Endpoint:     github.Endpoint,
	}

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Title": "rcmndr",
		}, "layouts/main")
	})

	app.Get("/auth/github/login", func(c *fiber.Ctx) error {
		authURL := oauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline)
		return c.Redirect(authURL)
	})

	app.Get("/auth/github/callback", func(c *fiber.Ctx) error {
		code := c.Query("code")
		token, err := oauth2Config.Exchange(c.Context(), code)
		// ...

		user, err := getUserInfo(token, c, oauth2Config)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		return c.SendString("Welcome " + user.Login)
	})
	// protected rout
	app.Get("/protected", func(c *fiber.Ctx) error {
		return c.SendString("protected")
	})

	log.Fatal(app.Listen(":3000"))
}
