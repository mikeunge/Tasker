package repository

import (
	"fmt"
	"os"
	"time"

	"github.com/mikeunge/Tasker/api/database"
	"github.com/mikeunge/Tasker/api/errors"

	"github.com/mikeunge/Tasker/api/entity"
	"gorm.io/gorm"
)

type taskRepository struct {
	entity.TaskRepository
	db *gorm.DB
}

func NewTaskRepository() *taskRepository {
	return &taskRepository{db: database.DB}
}

func (r *taskRepository) GetAll() ([]entity.Task, error) {
	var tasks []entity.Task
	res := r.db.Find(&tasks)
	if res.RowsAffected > 0 {
		return tasks, nil
	}
	return tasks, errors.NoTasksFound
}

func (r *taskRepository) Add(title, text string) (*entity.Task, error) {
	task, err := entity.NewTask(title, text)
	if err != nil {
		fmt.Println(err)
		return task, errors.FailedToAddTask
	}
	res := r.db.Create(&task)
	if res.Error != nil {
		return task, res.Error
	}
	return task, nil
}

func (r *taskRepository) Get(id string) (*entity.Task, error) {
	var task entity.Task
	res := r.db.Where("id = ?", id).First(&task)
	if res.Error != nil || res.RowsAffected == 0 {
		if res.Error != nil {
			fmt.Println(res.Error)
		}
		return &task, errors.TaskNotFound
	}
	return &task, nil
}

func (r *taskRepository) Update(id, title, text string) error {
	res := r.db.Model(&entity.Task{}).Where("id = ?", id).Updates(&entity.Task{Title: title, Text: text, UpdatedAt: time.Now()})
	if res.Error != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v", res.Error.Error())
		return errors.FailedToUpdateTask
	}
	return nil
}

// TODO: add the rest of the entity interfaces
