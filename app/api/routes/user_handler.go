package routes

import (
	"fmt"
	middleware2 "github.com/knands42/lorecrafter/app/api/middleware"
	"net/http"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// Me describe user info
// @Summary Get info about the logged user
// @Description Based on the logged user, get information from the token
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} string "User information"
// @Failure 401 {object} utils.ErrorResponse "Missing or invalid authorization header"
// @Router /api/me [get]
func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) error {
	userID := r.Context().Value(middleware2.UserIDContextKey)
	_, err := w.Write([]byte(fmt.Sprintf("Authenticated! User ID: %v", userID)))
	return err
}
