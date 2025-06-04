package domain

import (
	"errors"
	"github.com/knands42/lorecrafter/internal/utils"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrUUIDGeneration = errors.New("failed to generate UUID")
)

type UserCreationInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *UserCreationInput) Validate() error {
	var validationErrors []string

	if strings.TrimSpace(user.Username) == "" {
		validationErrors = append(validationErrors, "username is required")
	}

	if strings.TrimSpace(user.Email) == "" {
		validationErrors = append(validationErrors, "email is required")
	} else if !strings.Contains(user.Email, "@") {
		validationErrors = append(validationErrors, "email is invalid")
	}

	if len(user.Password) < 8 {
		validationErrors = append(validationErrors, "password must be at least 8 characters")
	}

	if len(validationErrors) > 0 {
		return &utils.ValidationError{Errors: validationErrors}
	}

	return nil
}

// User represents a user domain
type User struct {
	ID             uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	Active         bool      `json:"active"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}

// NewUser creates a new user with the given details
func NewUser(username, email, hashedPassword string) (*User, error) {
	now := time.Now()
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return nil, ErrUUIDGeneration
	}

	return &User{
		ID:             uuidV7,
		Username:       username,
		Email:          email,
		HashedPassword: hashedPassword,
		Active:         false,
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

func ToDomainUser(dbUser sqlc.User) (User, error) {
	id, err := utils.FromPGTypeUUID(dbUser.ID)
	if err != nil {
		return User{}, err
	}

	createdAt := dbUser.CreatedAt.Time
	updatedAt := dbUser.UpdatedAt.Time

	return User{
		ID:             id,
		Username:       dbUser.Username,
		Email:          dbUser.Email,
		HashedPassword: dbUser.HashedPassword,
		Active:         dbUser.IsActive,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}, nil
}
