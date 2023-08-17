package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
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

// session store
var store = session.New()

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve values
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	callbackURL := os.Getenv("GITHUB_CALLBACK_URL")

	// Create a new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/public", "./public")
	app.Use(printSessionMiddleware)

	oauth2Config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  callbackURL,
		Scopes:       []string{}, // Add required scopes here
		Endpoint:     github.Endpoint,
	}

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Title": "rcmndr",
		}, "layouts/main")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		authURL := oauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline)
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
		token, err := oauth2Config.Exchange(c.Context(), code)

		user, err := getUserInfo(token, c, oauth2Config)
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

func printSessionMiddleware(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		log.Println("Error retrieving session:", err)
		return c.Next()
	}

	// Print the entire session store
	log.Println("Session content:", sess.Fresh())

	// Or print specific values
	log.Println("User ID:", sess.Get("user_id"))

	return c.Next()
}
