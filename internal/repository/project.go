package repository

import (
	"fmt"

	"github.com/mikeunge/Tasker/internal/database"
	"github.com/mikeunge/Tasker/internal/entities"
	log "github.com/mikeunge/Tasker/pkg/logger"
)

// simple mapper for turnin DB projects into normal project structs
func transformDbObject(dbProject entities.IDBProject) entities.IProject {
	var project entities.IProject
	project.Id = dbProject.Id
	project.Name = dbProject.Name

	return project
}

func FindAllProjects() ([]entities.IProject, error) {
	var projects []entities.IProject

	db := database.Connection()
	query := "SELECT * FROM projects;"

	rows, err := db.Query(query)
	if err != nil {
		log.Error(err.Error())
		return projects, err
	}
	defer rows.Close()

	for rows.Next() {
		var dbProject entities.IDBProject
		if err := rows.Scan(&dbProject.Id, &dbProject.Name, &dbProject.CreatedAt); err != nil {
			return projects, err
		}
		projects = append(projects, transformDbObject(dbProject))
	}

	if err = rows.Err(); err != nil {
		return projects, err
	}

	return projects, nil
}

func FindProjectById(projectId int) (entities.IProject, error) {
	var project entities.IDBProject

	db := database.Connection()
	query := "SELECT * FROM projects WHERE id = $1;"

	err := db.QueryRow(query, projectId).Scan(&project.Id, &project.Name, &project.CreatedAt)
	if err != nil {
		return entities.IProject{}, err
	}

	return transformDbObject(project), nil
}

func UpdateOrCreateProject(project entities.IProject) (entities.IProject, error) {
	var p entities.IProject

	db := database.Connection()
	// query := fmt.Sprintf("INSERT OR REPLACE INTO projects (id, name) VALUES (%d, '%s') ON DUPLICATE KEY UPDATE name = '%s';", project.Id, project.Name, project.Name)
	query := fmt.Sprintf("INSERT INTO projects (id, name) VALUES (%d, '%s') ON CONFLICT (id) DO UPDATE SET name = '%s';", project.Id, project.Name, project.Name)
	stmt, err := db.Prepare(query)
	if err != nil {
		return p, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return p, err
	}
	return p, nil
}
