package middleware

import (
	"context"
	"github.com/knands42/lorecrafter/app/api/utils"
	"github.com/knands42/lorecrafter/internal/usecases"
	"log"
	"net/http"
	"strings"
)

const UserIDContextKey = "user_id"

func AuthMiddleware(authUseCase *usecases.AuthUseCase) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var tokenStr string

			// Check for Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader != "" {
				tokenStr = strings.TrimPrefix(authHeader, "Bearer ")
				if tokenStr == authHeader {
					utils.WriteJSONError(w, http.StatusUnauthorized, "malformed token")
					return
				}
			} else {
				// Check for auth_token cookie if header is not present
				cookie, err := r.Cookie("auth_token")
				log.Printf("Cookies: %s", r.Cookies())
				if err != nil || cookie.Value == "" {
					utils.WriteJSONError(w, http.StatusUnauthorized, "missing authentication credentials")
					return
				}
				tokenStr = cookie.Value
			}

			token, err := authUseCase.VerifyToken(tokenStr)
			if err != nil {
				utils.WriteJSONError(w, http.StatusUnauthorized, "invalid token")
				return
			}

			// Add user ID to context
			ctx := context.WithValue(r.Context(), UserIDContextKey, token.ID.String())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
