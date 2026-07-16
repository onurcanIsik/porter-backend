package models

import (
	"time"

	"github.com/google/uuid"
)

type ProjectModel struct {
	ID               uuid.UUID `db:"id" json:"id"`
	UserID           uuid.UUID `db:"user_id" json:"user_id"`
	ProjectName      string    `db:"project_name" json:"name"`
	ApiKey           string    `db:"api_key" json:"api_key"`
	ProjectCreatedAt time.Time `db:"project_created_at" json:"created_at"`
	ProjectUpdatedAt time.Time `db:"project_updated_at" json:"updated_at"`
}

type ProjectModelRepository interface {
	CreateProject(userID uuid.UUID, projectName, apiKey string) (*ProjectModel, error)
	GetProjectsByUserID(userID uuid.UUID) ([]*ProjectModel, error)
	GetProjectByID(projectID, userID uuid.UUID) (*ProjectModel, error)
	UpdateProject(projectID, userID uuid.UUID, projectName string) error
	DeleteProject(projectID, userID uuid.UUID) error
}

type CreateProjectRequest struct {
	ProjectName string `json:"project_name"`
}

type UpdateProjectRequest struct {
	ProjectName string `json:"project_name"`
}
