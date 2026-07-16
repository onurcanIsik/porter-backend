package main

import (
	"log"
	"net/http"
	"os"
	"porter/internal/handler"
	"porter/internal/repo"
	"porter/internal/service"
	"porter/middleware"
	"porter/pkg/database"
	"porter/pkg/jwt"

	"github.com/joho/godotenv"
)

func main() {

	database, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	if googleClientID == "" {
		log.Fatal("GOOGLE_CLIENT_ID is not set in the environment variables")
	}

	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	if googleClientSecret == "" {
		log.Fatal("GOOGLE_CLIENT_SECRET is not set in the environment variables")
	}

	jwtManager := jwt.NewJWTManager(os.Getenv("SECRET_KEY"))

	// REPOS
	userRepo := repo.NewUserRepo(database)
	refreshTokenRepo := repo.NewRefreshTokenRepo(database)
	quotaRepo := repo.NewQuotaRepo(database)
	projectsRepo := repo.NewProjectsRepo(database)

	// SERVICES
	userService := service.NewUserService(userRepo, jwtManager, refreshTokenRepo)
	quotaService := service.NewQuotaService(quotaRepo)
	projectsService := service.NewProjectsService(projectsRepo)

	// HANDLERS
	userHandler := handler.NewUserHandler(userService, jwtManager)
	quotaHandler := handler.NewQuotaHandler(quotaService)
	projectsHandler := handler.NewProjectsHandler(projectsService)

	// ROUTERS

	mux := http.NewServeMux()

	rateLimiter := middleware.NewRateLimiter()
	requireAuth := middleware.RequireAuth(jwtManager)
	wrappedMux := rateLimiter.Middleware(mux)

	// AUTH
	mux.HandleFunc("/api/auth/google", userHandler.GoogleLogin)
	mux.HandleFunc("/api/auth/google_callback", userHandler.GoogleCallback)

	// REFRESH TOKEN
	mux.HandleFunc("POST /api/refresh", userHandler.RefreshToken)

	// QUOTA
	mux.Handle("GET /api/quota", requireAuth(http.HandlerFunc(quotaHandler.GetQuotaByUserID)))
	mux.Handle("PUT /api/quota/update", requireAuth(http.HandlerFunc(quotaHandler.UpdateQuota)))

	// PROJECTS
	mux.Handle("GET /api/projects", requireAuth(http.HandlerFunc(projectsHandler.GetProjectsService)))
	mux.Handle("POST /api/projects", requireAuth(http.HandlerFunc(projectsHandler.CreateProject)))
	mux.Handle("GET /api/projects/{id}", requireAuth(http.HandlerFunc(projectsHandler.GetProjectByID)))
	mux.Handle("PUT /api/projects/{id}", requireAuth(http.HandlerFunc(projectsHandler.UpdateProject)))
	mux.Handle("DELETE /api/projects/{id}", requireAuth(http.HandlerFunc(projectsHandler.DeleteProject)))

	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
