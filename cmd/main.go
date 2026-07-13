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

	// SERVICES
	userService := service.NewUserService(userRepo, jwtManager, refreshTokenRepo)
	quotaService := service.NewQuotaService(quotaRepo)

	// HANDLERS
	userHandler := handler.NewUserHandler(userService, jwtManager)
	quotaHandler := handler.NewQuotaHandler(quotaService)

	// ROUTERS

	mux := http.NewServeMux()

	rateLimiter := middleware.NewRateLimiter()
	wrappedMux := rateLimiter.Middleware(mux)

	// AUTH
	mux.HandleFunc("/api/auth/google", userHandler.GoogleLogin)
	mux.HandleFunc("/api/auth/google_callback", userHandler.GoogleCallback)

	// REFRESH TOKEN
	mux.HandleFunc("POST /api/refresh", userHandler.RefreshToken)

	// QUOTA
	mux.HandleFunc("/api/quota", quotaHandler.GetQuotaByUserID)
	mux.HandleFunc("/api/quota/update", quotaHandler.UpdateQuota)

	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
