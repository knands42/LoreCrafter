package domain

import (
	"github.com/knands42/lorecrafter/internal/utils"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"time"
)

// LoginInput represents a request to authenticate a user
type LoginInput struct {
	Username string
	Password string
}

func (input *LoginInput) Validate() error {
	var validationErrors []string

	if len(input.Username) < 4 {
		validationErrors = append(validationErrors, "username must be at least 5 characters long")
	}

	if len(input.Password) < 8 {
		validationErrors = append(validationErrors, "password must be at least 8 characters long")
	}

	if len(validationErrors) > 0 {
		return &utils.ValidationError{Errors: validationErrors}
	}

	return nil
}

// AuthOutput represents the response after successful authentication
type AuthOutput struct {
	User      sqlc.User
	Token     string
	ExpiresAt time.Time
}

// TokenPayload represents the data stored in the authentication token
type TokenPayload struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
