package repo

import (
	"context"
	"porter/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectsRepo struct {
	db *pgxpool.Pool
}

func NewProjectsRepo(db *pgxpool.Pool) *ProjectsRepo {
	return &ProjectsRepo{db: db}
}

func (r *ProjectsRepo) CreateProject(userID uuid.UUID, projectName, apiKey string) (*models.ProjectModel, error) {

	query := `INSERT INTO projects (user_id, project_name, api_key)
	          VALUES ($1, $2, $3)
	          RETURNING id, project_created_at, project_updated_at`

	project := models.ProjectModel{
		UserID:      userID,
		ProjectName: projectName,
		ApiKey:      apiKey,
	}

	err := r.db.QueryRow(context.Background(), query,
		userID, projectName, apiKey,
	).Scan(&project.ID, &project.ProjectCreatedAt, &project.ProjectUpdatedAt)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (r *ProjectsRepo) GetProjectsByUserID(userID uuid.UUID) ([]*models.ProjectModel, error) {

	query := `SELECT id, user_id, project_name, api_key, project_created_at, project_updated_at FROM projects WHERE user_id = $1`

	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*models.ProjectModel

	for rows.Next() {
		var project models.ProjectModel
		err := rows.Scan(&project.ID, &project.UserID, &project.ProjectName, &project.ApiKey, &project.ProjectCreatedAt, &project.ProjectUpdatedAt)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *ProjectsRepo) GetProjectByID(projectID, userID uuid.UUID) (*models.ProjectModel, error) {
	query := `SELECT id, user_id, project_name, api_key, project_created_at, project_updated_at FROM projects WHERE id = $1 AND user_id = $2`
	var project models.ProjectModel
	err := r.db.QueryRow(context.Background(), query, projectID, userID).Scan(
		&project.ID,
		&project.UserID,
		&project.ProjectName,
		&project.ApiKey,
		&project.ProjectCreatedAt,
		&project.ProjectUpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *ProjectsRepo) UpdateProject(projectID, userID uuid.UUID, projectName string) error {
	query := `UPDATE projects SET project_name = $1, project_updated_at = NOW() WHERE id = $2 AND user_id = $3`
	_, err := r.db.Exec(context.Background(), query, projectName, projectID, userID)
	return err
}

func (r *ProjectsRepo) DeleteProject(projectID, userID uuid.UUID) error {
	query := `DELETE FROM projects WHERE id = $1 AND user_id = $2`
	_, err := r.db.Exec(context.Background(), query, projectID, userID)
	return err
}
