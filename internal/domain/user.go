package domain

import (
	"errors"
	"github.com/knands42/lorecrafter/internal/utils"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"strings"
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

func (user *UserCreationInput) ToSqlcParams(hashedPassword string) (sqlc.CreateUserParams, error) {
	newUUUIDV7, err := utils.GeneratePGUUID()
	if err != nil {
		return sqlc.CreateUserParams{}, err
	}

	return sqlc.CreateUserParams{
		ID:             newUUUIDV7,
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: hashedPassword,
	}, nil
}
