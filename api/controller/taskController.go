package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mikeunge/Tasker/api/repository"
	"github.com/mikeunge/Tasker/api/utils"
)

func Register(app *fiber.App) {
	api := app.Group("/api/tasks")
	api.Get("/", getAllTasks)
	api.Post("/", addTask)
	api.Get("/:id?", getTaskById)
	api.Put("/:id?", updateTask)
	api.Delete("/:id?", deleteTask)
}

func getAllTasks(c *fiber.Ctx) error {
	taskRepo := repository.NewTaskRepository()
	tasks, err := taskRepo.GetAll()
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data":  make([]string, 0),
			"error": fmt.Sprintf("%+v", err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  tasks,
		"error": "",
	})
}

func getTaskById(c *fiber.Ctx) error {
	id := c.Params("id")
	if !utils.IsValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":  make([]string, 0),
			"error": "provided id is not a valid id",
		})
	}
	taskRepo := repository.NewTaskRepository()
	task, err := taskRepo.Get(id)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":  make([]string, 0),
			"error": fmt.Sprintf("%+v", err),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":  task,
		"error": "",
	})
}

func addTask(c *fiber.Ctx) error {
	taskRepo := repository.NewTaskRepository()
	task, err := taskRepo.Add("abc", "dieser text")
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":  make([]string, 0),
			"error": fmt.Sprintf("%+v", err),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":  task,
		"error": "",
	})
}

func updateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if !utils.IsValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":  make([]string, 0),
			"error": "provided id is not a valid id",
		})
	}
	taskRepo := repository.NewTaskRepository()
	err := taskRepo.Update(id, "cbd", "neuer text")
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":  make([]string, 0),
			"error": fmt.Sprintf("%+v", err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  "success",
		"error": "",
	})
}

func deleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if !utils.IsValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":  make([]string, 0),
			"error": "provided id is not a valid id",
		})
	}
	taskRepo := repository.NewTaskRepository()
	err := taskRepo.Update(id, "cbd", "neuer text")
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":  make([]string, 0),
			"error": fmt.Sprintf("%+v", err),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  "success",
		"error": "",
	})
}
