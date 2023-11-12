package repository

import (
	"fmt"

	"github.com/mikeunge/Tasker/internal/database"
	"github.com/mikeunge/Tasker/internal/entities"
	log "github.com/mikeunge/Tasker/pkg/logger"
)

type IProjectRepository interface {
	FindAllProjects() ([]entities.IProject, error)
	FindProjectById(projectId int) (entities.IProject, error)
	CreateProject(name string) (int, error)
	CreateAndReturnProject(name string) (entities.IProject, error)
}

type ProjectRepository struct {
	Project *entities.IProject
}

func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{
		Project: new(entities.IProject),
	}
}

func (repo *ProjectRepository) FindAllProjects() ([]entities.IProject, error) {
	var projects []entities.IProject
	db := database.Connection()
	query := "SELECT * FROM projects;"

	rows, err := db.Query(query)
	if err != nil {
		log.Error("Could not fetch projects, Error: %+v", err)
		return projects, err
	}
	defer rows.Close()

	for rows.Next() {
		var p entities.IProject
		if err := rows.Scan(&p.Id, &p.Name, &p.CreatedAt); err != nil {
			return projects, err
		}
		projects = append(projects, p)
	}

	if err = rows.Err(); err != nil {
		return projects, err
	}

	return projects, nil
}

func (repo ProjectRepository) FindProjectById(projectId int) (entities.IProject, error) {
	var project entities.IProject
	db := database.Connection()
	query := fmt.Sprintf("SELECT * FROM projects WHERE id = '%d';", projectId)
	err := db.QueryRow(query).Scan(&project.Id, &project.Name, &project.CreatedAt)
	if err != nil {
		return project, err
	}

	return project, nil
}

func (repo ProjectRepository) CreateProject(name string) (int, error) {
	var projectId int64
	db := database.Connection()
	query := fmt.Sprintf("INSERT INTO projects (name) VALUES ('%s');", name)
	stmt, err := db.Prepare(query)
	if err != nil {
		return int(projectId), err
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		return int(projectId), err
	}

	projectId, err = res.LastInsertId()
	if err != nil {
		return int(projectId), err
	}

	return int(projectId), err
}

func (repo ProjectRepository) CreateAndReturnProject(name string) (entities.IProject, error) {
	var project entities.IProject

	id, err := repo.CreateProject(name)
	if err != nil {
		return project, err
	}

	project, err = repo.FindProjectById(int(id))
	if err != nil {
		return project, err
	}

	return project, nil
}
