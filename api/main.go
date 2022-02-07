package main

import (
	"fmt"
	"os"

	"github.com/mikeunge/Tasker/api/database"
	"github.com/mikeunge/Tasker/api/repository"
	"github.com/mikeunge/Tasker/api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	db, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not connect to database.\n\nError: %v\n", err)
		os.Exit(1)
	}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		taskRepo := repository.NewTaskRepository(db)
		tasks, err := taskRepo.GetAll()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%+v", tasks)
		}
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/add", func(c *fiber.Ctx) error {
		taskRepo := repository.NewTaskRepository(db)
		task, err := taskRepo.Add()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Print(task)
		}
		return c.SendString("Create task")
	})

	port, err := utils.GetEnv("APP_PORT")
	if err != nil {
		fmt.Fprintf(os.Stderr, "APP_PORT not specified or not able to load from .env, using fallback :: 3030")
		port = "3030"
	}

	app.Listen(":" + port)
}
