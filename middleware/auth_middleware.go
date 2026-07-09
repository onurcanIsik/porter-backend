package middleware

import (
	"context"
	"net/http"
	"porter/pkg/jwt"
	"strings"
)

type contextKey string

const userIDKey contextKey = "user_id"

func RequireAuth(manager *jwt.JWTManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userID, err := manager.ValidateToken(tokenString, "access")
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, userIDKey, userID)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)

		})
	}
}
