package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/knands42/lorecrafter/app/api/middleware"
	"github.com/knands42/lorecrafter/app/api/utils"

	"github.com/go-chi/chi/v5"
	"github.com/knands42/lorecrafter/internal/domain"
	"github.com/knands42/lorecrafter/internal/usecases"
)

// AuthHandler handles authentication-related HTTP requests
type AuthHandler struct {
	authUseCase          *usecases.AuthUseCase
	passwordResetUseCase *usecases.PasswordResetUseCase
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(authUseCase *usecases.AuthUseCase, passwordResetUseCase *usecases.PasswordResetUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase:          authUseCase,
		passwordResetUseCase: passwordResetUseCase,
	}
}

// RegisterRoutes registers the authentication routes
func (h *AuthHandler) RegisterRoutes(r chi.Router) {
	r.Post("/register", middleware.ErrorHandlerMiddleware(h.Register))
	r.Post("/login", middleware.ErrorHandlerMiddleware(h.Login))
	r.Post("/logout", middleware.ErrorHandlerMiddleware(h.Logout))

	// Password reset endpoints
	r.Post("/forgot-password", middleware.ErrorHandlerMiddleware(h.ForgotPassword))
	r.Post("/reset-password", middleware.ErrorHandlerMiddleware(h.ResetPassword))
}

// Register handles user registration
// @Summary Register a new user
// @Description Register a new user with the provided credentials
// @Tags auth
// @Accept json
// @Produce json
// @Param input body domain.UserCreationInput true "User registration details"
// @Success 201 {object} domain.User "User registered successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request body"
// @Failure 409 {object} utils.ErrorResponse "User already exists"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/auth/register [post]
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
// @Summary Login a user
// @Description Login a user with the provided credentials
// @Tags auth
// @Accept json
// @Produce json
// @Param input body domain.LoginInput true "User login details"
// @Success 200 {object} domain.AuthOutput "User logged in successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request body"
// @Failure 401 {object} utils.ErrorResponse "Invalid credentials"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) error {
	var input domain.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
	}

	response, err := h.authUseCase.Login(input)
	if err != nil {
		switch {
		case errors.Is(err, usecases.ErrInvalidCredentials):
			return utils.WriteJSONError(w, http.StatusUnauthorized, "Invalid credentials")
		case errors.Is(err, usecases.ErrUserNotFound):
			return utils.WriteJSONError(w, http.StatusUnauthorized, "Invalid credentials")
		}

		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    response.Token,
		Path:     "/",
		HttpOnly: true,
		Secure:   r.TLS != nil,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   60 * 60 * 24, // 24 hours
	})

	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(response)
}

// Logout clears the authentication token by setting an empty cookie with the same name.
// @Summary Logout a user
// @Description Logout a user by clearing the authentication token
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 204 "User logged out successfully"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) error {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   r.TLS != nil,
		SameSite: http.SameSiteLaxMode,
	})

	return nil
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

// ForgotPassword handles the forgot password request
// @Summary Forgot password
// @Description Forgot password
// @Tags auth
// @Accept json
// @Produce json
// @Param input body domain.ForgotPasswordInput true "User forgot password details"
// @Success 204 "Password reset email sent successfully"
// @Failure 400 {object} utils.ErrorResponse "Failed to generate token"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/auth/forgot-password [post]
func (h *AuthHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) error {
	var input domain.ForgotPasswordInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
	}

	err := h.passwordResetUseCase.RequestPasswordReset(input)
	if err != nil {
		switch {
		case errors.Is(err, usecases.ErrFailedToCreatePasswordResetToken):
			return utils.WriteJSONError(w, http.StatusBadRequest, usecases.ErrFailedToCreatePasswordResetToken.Error())
		case errors.Is(err, usecases.ErrFailedToGenerateToken):
			return utils.WriteJSONError(w, http.StatusBadRequest, usecases.ErrFailedToGenerateToken.Error())
		}

		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// ResetPassword handles the reset password request
// @Summary Reset password
// @Description Reset password
// @Tags auth
// @Accept json
// @Produce json
// @Param input body domain.PasswordResetInput true "User reset password details"
// @Success 204 "Password reset successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request body"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /api/auth/reset-password [post]
func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) error {
	var input domain.PasswordResetInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
	}

	err := h.passwordResetUseCase.ResetPassword(input)
	if err != nil {
		switch {
		case errors.Is(err, usecases.ErrInvalidResetToken):
			return utils.WriteJSONError(w, http.StatusBadRequest, usecases.ErrInvalidResetToken.Error())
		case errors.Is(err, usecases.ErrPasswordResetTokenExpired):
			return utils.WriteJSONError(w, http.StatusBadRequest, usecases.ErrPasswordResetTokenExpired.Error())
		case errors.Is(err, usecases.ErrPasswordResetTokenUsed):
			return utils.WriteJSONError(w, http.StatusBadRequest, usecases.ErrPasswordResetTokenUsed.Error())
		}

		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
