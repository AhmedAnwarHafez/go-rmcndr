package utils

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type GitHubUser struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
}

func GetUserInfo(token *oauth2.Token, ctx *fiber.Ctx, config oauth2.Config) (*GitHubUser, error) {
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

func GetConfig() oauth2.Config {
	// Retrieve values
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	callbackURL := os.Getenv("GITHUB_CALLBACK_URL")

	oauth2Config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  callbackURL,
		Scopes:       []string{}, // Add required scopes here
		Endpoint:     github.Endpoint,
	}

	return oauth2Config
}
