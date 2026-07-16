package service

import (
	"porter/models"
	apprand "porter/pkg/crypto"

	"github.com/google/uuid"
)

type ProjectsService struct {
	projectRepo models.ProjectModelRepository
}

func NewProjectsService(projectRepo models.ProjectModelRepository) *ProjectsService {
	return &ProjectsService{
		projectRepo: projectRepo,
	}
}

func (s *ProjectsService) CreateProject(userID uuid.UUID, projectName string) (*models.ProjectModel, error) {
	urlApiKey, err := apprand.GenerateRandomSafeString(16)
	if err != nil {
		return nil, err
	}

	return s.projectRepo.CreateProject(userID, projectName, urlApiKey)
}

func (s *ProjectsService) GetProjectsByUserID(userID uuid.UUID) ([]*models.ProjectModel, error) {

	return s.projectRepo.GetProjectsByUserID(userID)
}

func (s *ProjectsService) GetProjectByID(projectID, userID uuid.UUID) (*models.ProjectModel, error) {
	return s.projectRepo.GetProjectByID(projectID, userID)
}

func (s *ProjectsService) UpdateProject(projectID, userID uuid.UUID, projectName string) error {
	return s.projectRepo.UpdateProject(projectID, userID, projectName)
}

func (s *ProjectsService) DeleteProject(projectID, userID uuid.UUID) error {
	return s.projectRepo.DeleteProject(projectID, userID)
}
