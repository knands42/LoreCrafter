package routes

import (
	"encoding/json"
	middleware2 "github.com/knands42/lorecrafter/app/api/middleware"
	"net/http"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

type AboutMe struct {
	UserId string `json:"user_id"`
}

// Me describe user info
// @Summary Get info about the logged user
// @Description Based on the logged user, get information from the token
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} AboutMe "User Info"
// @Failure 401 {object} utils.ErrorResponse "Missing or invalid authorization header"
// @Router /api/me [get]
func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) error {
	userID := r.Context().Value(middleware2.UserIDContextKey)
	castUserID, _ := userID.(string)

	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(AboutMe{
		UserId: castUserID,
	})
}
