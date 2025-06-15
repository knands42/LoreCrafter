package middleware

import (
	"context"
	"github.com/knands42/lorecrafter/app/api/utils"
	"github.com/knands42/lorecrafter/internal/usecases"
	"net/http"
	"strings"
)

const UserIDContextKey = "user_id"

func AuthMiddleware(authUseCase *usecases.AuthUseCase) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.WriteJSONError(w, http.StatusUnauthorized, "missing authorization header")
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenStr == authHeader {
				utils.WriteJSONError(w, http.StatusUnauthorized, "malformed token")
				return
			}

			token, err := authUseCase.VerifyToken(tokenStr)
			if err != nil {
				utils.WriteJSONError(w, http.StatusUnauthorized, "invalid token")
				return
			}

			// Add user ID to context
			ctx := context.WithValue(r.Context(), UserIDContextKey, token.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
