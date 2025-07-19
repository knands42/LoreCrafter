package routes

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	middleware2 "github.com/knands42/lorecrafter/app/api/middleware"
	"github.com/knands42/lorecrafter/app/api/utils"
	"github.com/knands42/lorecrafter/internal/usecases"
	"net/http"
)

type UserHandler struct {
	userUseCase *usecases.UserUseCase
}

func NewUserHandler(
	userUseCase *usecases.UserUseCase,
) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// Me describe user info
// @Summary Get info about the logged user
// @Description Based on the logged user, get information from the token
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.User "User Info"
// @Failure 401 {object} utils.ErrorResponse "Missing or invalid authorization header"
// @Failure 404 {object} utils.ErrorResponse "User not found"
// @Router /api/me [get]
func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) error {
	userIDStr, ok := r.Context().Value(middleware2.UserIDContextKey).(string)
	if !ok {
		return utils.WriteJSONError(w, http.StatusUnauthorized, "User ID not found in context")
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
	}

	userInfo, err := h.userUseCase.GetUserInfo(userID)
	if err != nil {
		switch {
		case errors.Is(err, usecases.ErrUserNotFound):
			return utils.WriteJSONError(w, http.StatusNotFound, err.Error())
		}
	}
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(userInfo)
}
