package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	app.Post("/hello", func(c *fiber.Ctx) error {
		return c.
	})
	app.Listen(":8080")
}
