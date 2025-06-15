package domain

import (
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"time"
)

// LoginInput represents a request to authenticate a user
type LoginInput struct {
	Username string
	Password string
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
