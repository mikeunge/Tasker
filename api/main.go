package main

import (
	"tasker-api/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	port, err := utils.GetEnv("APP_PORT")
	if err != nil {
		port = "3033"
	}

	app.Listen(":" + port)
}
