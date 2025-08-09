package middleware

import (
	"context"
	"net/http"

	"github.com/atbu/slate/backend/auth"
	"github.com/google/uuid"
)

type contextKey string

const (
	UserIDKey contextKey = "userID"
)

func AuthMiddleware(authService *auth.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("auth_token")
			if err != nil || cookie == nil {
				http.Error(w, "Not logged in", http.StatusUnauthorized)
				return
			}

			claims, err := authService.ValidateToken(cookie.Value)
			if err != nil {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			userIDStr, ok := claims["sub"].(string)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			userID, err := uuid.Parse(userIDStr)
			if err != nil {
				http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserID(r *http.Request) (uuid.UUID, bool) {
	userID, ok := r.Context().Value(UserIDKey).(uuid.UUID)
	return userID, ok
}
