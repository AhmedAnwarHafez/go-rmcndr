package main

import (
	"github.com/gofiber/fiber/v2"
)

func GetSearchHandler(c *fiber.Ctx) error {

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
}
