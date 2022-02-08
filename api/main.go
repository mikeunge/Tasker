package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mikeunge/Tasker/api/controller"
	"github.com/mikeunge/Tasker/api/database"
	"github.com/mikeunge/Tasker/api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
	tz, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Setting the timezone went wrong, exiting.\n\nError: %v\n", err)
		os.Exit(1)
	}
	time.Local = tz
}

func main() {
	err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not connect to database.\n\nError: %v\n", err)
		os.Exit(1)
	}

	app := fiber.New()
	app.Use(logger.New())

	controller.Register(app)

	port, err := utils.GetEnv("APP_PORT")
	if err != nil {
		fmt.Fprintf(os.Stderr, "APP_PORT not specified or not able to load from .env, using fallback :: 3030")
		port = "3030"
	}

	app.Listen(":" + port)
}
