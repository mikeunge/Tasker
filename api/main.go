package main

import (
	"fmt"
	"os"
	"tasker-api/database"
	"tasker-api/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	_, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not connect to database.\n\nError: %v\n", err)
		os.Exit(1)
	}

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
