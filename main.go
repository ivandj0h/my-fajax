package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	// Setting views engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})
}
