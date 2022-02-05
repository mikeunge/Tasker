package main

import (
	"fmt"
	"os"
	"tasker-api/database"
	"tasker-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
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
		fmt.Fprintf(os.Stderr, "APP_PORT not specified or not able to load from .env, using fallback :: 3030")
		port = "3030"
	}

	app.Listen(":" + port)
}
