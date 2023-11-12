package repository

import (
	"fmt"

	"github.com/mikeunge/Tasker/internal/database"
	"github.com/mikeunge/Tasker/internal/entities"
	log "github.com/mikeunge/Tasker/pkg/logger"
)

type ITaskRepository interface {
	FindAllTasksForProject(projectId int) ([]entities.ITask, error)
	FindTaskById(taskId int) (entities.ITask, error)
	CreateTask(projectId int, title string, description string, statusId int) (int, error)
	CreateAndReturnTask(projectId int, title string, description string, statusId int) (entities.ITask, error)
}

type TaskRepository struct {
	Tasks *[]entities.ITask
}

func NewTaskRepository() *TaskRepository {
	tasks := make([]entities.ITask, 0)
	return &TaskRepository{
		Tasks: &tasks,
	}
}

func (repo *TaskRepository) FindAllTasksForProject(projectId int) ([]entities.ITask, error) {
	var tasks []entities.ITask
	var db = database.Connection()
	var query string = fmt.Sprintf("SELECT * FROM tasks WHERE project_id = '%d';", projectId)

	rows, err := db.Query(query)
	if err != nil {
		log.Error("Could not fetch projects, Error: %+v", err)
		return tasks, err
	}
	defer rows.Close()

	for rows.Next() {
		var task entities.ITask
		if err := rows.Scan(
			&task.Id, &task.ProjectId, &task.Title,
			&task.Description, &task.StatusId, &task.CreatedAt,
		); err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (repo TaskRepository) FindTaskById(taskId int) (entities.ITask, error) {
	var task entities.ITask
	db := database.Connection()
	query := fmt.Sprintf("SELECT * FROM task WHERE id = '%d';", taskId)
	err := db.QueryRow(query).Scan(
		&task.Id, &task.ProjectId, &task.Title,
		&task.Description, &task.StatusId, &task.CreatedAt)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (repo TaskRepository) CreateTask(
	projectId int, title string, description string, statusId int,
) (int, error) {
	var taskId int64
	db := database.Connection()
	query := fmt.Sprintf(
		"INSERT INTO tasks ('project_id', 'title', 'description', 'status_id') VALUES ('%d', '%s', '%s', '%d');",
		projectId, title, description, statusId)

	stmt, err := db.Prepare(query)
	if err != nil {
		return int(taskId), err
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		return int(taskId), err
	}

	taskId, err = res.LastInsertId()
	if err != nil {
		return int(taskId), err
	}

	return int(taskId), err
}

func (repo TaskRepository) CreateAndReturnTask(
	projectId int, title string, description string, statusId int,
) (entities.ITask, error) {
	var task entities.ITask

	id, err := repo.CreateTask(projectId, title, description, statusId)
	if err != nil {
		return task, err
	}

	task, err = repo.FindTaskById(int(id))
	if err != nil {
		return task, err
	}

	return task, nil
}
