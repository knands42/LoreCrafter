package routes

import (
	"encoding/json"
	"errors"
	"github.com/knands42/lorecrafter/cmd/api/middleware"
	"github.com/knands42/lorecrafter/cmd/api/utils"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/knands42/lorecrafter/internal/domain"
	"github.com/knands42/lorecrafter/internal/usecases"
)

// AuthHandler handles authentication-related HTTP requests
type AuthHandler struct {
	authUseCase *usecases.AuthUseCase
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(authUseCase *usecases.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

// RegisterRoutes registers the authentication routes
func (h *AuthHandler) RegisterRoutes(r chi.Router) {
	r.Post("/register", middleware.ErrorHandlerMiddleware(h.Register))
	r.Post("/login", middleware.ErrorHandlerMiddleware(h.Login))
}

// Register handles user registration
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) error {
	var input domain.UserCreationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
	}

	response, err := h.authUseCase.Register(input)
	if err != nil {
		switch {
		case errors.Is(err, usecases.ErrUserAlreadyExists):
			return utils.WriteJSONError(w, http.StatusConflict, "UserCreation already exists")

		case errors.Is(err, usecases.ErrInvalidCredentials):
			return utils.WriteJSONError(w, http.StatusUnauthorized, "Invalid credentials")
		}

		return err
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(response)
}

// Login handles user login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) error {
	var input domain.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
	}

	response, err := h.authUseCase.Login(input)
	if err != nil {
		if errors.Is(err, usecases.ErrInvalidCredentials) {
			return utils.WriteJSONError(w, http.StatusUnauthorized, "Invalid credentials")
		}

		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

// GetAuthorizationPayload extracts the token payload from the Authorization header
func (h *AuthHandler) GetAuthorizationPayload(r *http.Request) (*domain.TokenPayload, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header is not provided")
	}

	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, errors.New("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != "bearer" {
		return nil, errors.New("unsupported authorization type")
	}

	token := fields[1]
	payload, err := h.authUseCase.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
