package services

import (
	"github.com/mikeunge/Tasker/internal/entities"
	"github.com/mikeunge/Tasker/internal/repository"
	log "github.com/mikeunge/Tasker/pkg/logger"
)

func NewProject(name string) (entities.IProject, error) {
	var project entities.IProject
	project.Name = name

	log.Debug("Creating new project: %s", name)
	newProject, err := repository.UpdateOrCreateProject(project)
	if err != nil {
		log.Error("Could not create project: %s; ERR: %+v", name, err)
		return project, err
	}

	return newProject, nil
}

func GetProjectById(name string) {}

func UpdateProjectName(project *entities.IProject, name string) {}
