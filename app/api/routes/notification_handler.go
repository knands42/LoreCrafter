package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/knands42/lorecrafter/app/api/middleware"
	"github.com/knands42/lorecrafter/app/api/utils"
	"github.com/knands42/lorecrafter/internal/usecases"
)

type NotificationHandler struct {
	notificationUseCase *usecases.NotificationUseCase
	authUseCase         *usecases.AuthUseCase
}

func NewNotificationHandler(
	notificationUseCase *usecases.NotificationUseCase,
	authUseCase *usecases.AuthUseCase,
) *NotificationHandler {
	return &NotificationHandler{
		notificationUseCase: notificationUseCase,
		authUseCase:         authUseCase,
	}
}

func (h *NotificationHandler) RegisterRoutes(r chi.Router) {
	r.Route("/notifications", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(h.authUseCase))

		// Get all notifications for the current user
		r.Get("/", middleware.ErrorHandlerMiddleware(h.GetNotifications))
	})
}

// GetNotifications returns all notifications for the current user
// @Summary Get notifications
// @Description Get all notifications for the current user
// @Tags notifications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} domain.Notification "Notifications retrieved successfully"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/notifications [get]
func (h *NotificationHandler) GetNotifications(w http.ResponseWriter, r *http.Request) error {
	// Get user ID from context (set by auth middleware)
	userIDStr, ok := r.Context().Value(middleware.UserIDContextKey).(string)
	if !ok {
		return utils.WriteJSONError(w, http.StatusUnauthorized, "User ID not found in context")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
	}

	// Get notifications from usecase
	notifications, err := h.notificationUseCase.GetNotifications(userID)
	if err != nil {
		return utils.WriteJSONError(w, http.StatusInternalServerError, "Failed to get notifications")
	}

	// Return notifications as JSON
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(notifications)
}
