package repository

import (
	"errors"

	"github.com/mikeunge/Tasker/api/entity"
	"gorm.io/gorm"
)

var (
	errorTaskNotFound       = errors.New("the task does not exist")
	errorNoTasksFound       = errors.New("there are no tasks in the db")
	errorFailedToAddTask    = errors.New("could not create new task")
	errorFailedToUpdateTask = errors.New("failed to update task")
)

type taskRepository struct {
	entity.TaskRepository
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) GetAll() ([]entity.Task, error) {
	var tasks []entity.Task
	result := r.db.Find(&tasks)
	if result.RowsAffected > 0 {
		return tasks, nil
	}
	return tasks, errorNoTasksFound
}

func (r *taskRepository) Add() (string, error) {
	task, err := entity.NewTask("title", "text")
	if err != nil {
		return "", err
	}
	res := r.db.Create(&task)
	if res.Error != nil {
		return "", res.Error
	}
	return task.Id, nil
}

// TODO: add the rest of the entity interfaces
