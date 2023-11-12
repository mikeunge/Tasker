package services

import (
	"github.com/mikeunge/Tasker/internal/entities"
	"github.com/mikeunge/Tasker/internal/repository"
	log "github.com/mikeunge/Tasker/pkg/logger"
)

var projectRepository *repository.ProjectRepository

func init() {
	projectRepository = repository.NewProjectRepository()
}

func NewProject(name string) (*entities.IProject, error) {
	var err error

	log.Debug("Creating new project: %s", name)
	*projectRepository.Project, err = projectRepository.CreateAndReturnProject(name)
	if err != nil {
		log.Error("Could not create project: %s; Error: %+v", name, err)
		return &entities.IProject{}, err
	}

	return projectRepository.Project, nil
}

func GetProjectById(name string) {}

func UpdateProjectName(project *entities.IProject, name string) {}
