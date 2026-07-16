package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"porter/internal/service"
	"porter/middleware"
	"porter/models"
	apperr "porter/pkg/err"
	"strings"

	"github.com/google/uuid"
)

type ProjectsHandler struct {
	projectsService *service.ProjectsService
}

func NewProjectsHandler(projectsService *service.ProjectsService) *ProjectsHandler {
	return &ProjectsHandler{
		projectsService: projectsService,
	}
}

func (h *ProjectsHandler) GetProjectsService(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	projects, err := h.projectsService.GetProjectsByUserID(userID)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(projects); err != nil {
		log.Printf("failed to encode projects response: %v", err)
		return
	}

}

func (h *ProjectsHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req models.CreateProjectRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if req.ProjectName == "" {
		http.Error(w, "project_name is required", http.StatusBadRequest)
		return
	}

	if len(req.ProjectName) > 100 {
		http.Error(w, "project_name must be less than 100 characters", http.StatusBadRequest)
		return
	}

	if len(req.ProjectName) < 3 {
		http.Error(w, "project_name must be at least 3 characters", http.StatusBadRequest)
		return
	}

	if !isValidProjectName(req.ProjectName) {
		http.Error(w, "project_name can only contain letters, numbers, hyphens, and underscores", http.StatusBadRequest)
		return
	}

	req.ProjectName = strings.TrimSpace(req.ProjectName)

	project, err := h.projectsService.CreateProject(userID, req.ProjectName)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(project); err != nil {
		log.Printf("failed to encode project response: %v", err)
		return
	}

}

func (h *ProjectsHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	projectID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	project, err := h.projectsService.GetProjectByID(projectID, userID)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(project); err != nil {
		log.Printf("failed to encode project response: %v", err)
		return
	}

}

func (h *ProjectsHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	projectID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	err = h.projectsService.DeleteProject(projectID, userID)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (h *ProjectsHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	projectID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	var req models.UpdateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	err = h.projectsService.UpdateProject(projectID, userID, req.ProjectName)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204: yaptım, anlatacak bir şeyim yok
}

func isValidProjectName(name string) bool {
	for _, char := range name {
		if !(char >= 'a' && char <= 'z') &&
			!(char >= 'A' && char <= 'Z') &&
			!(char >= '0' && char <= '9') &&
			char != '-' && char != '_' {
			return false
		}
	}
	return true
}
