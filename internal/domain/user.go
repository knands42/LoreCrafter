package domain

import (
	"github.com/google/uuid"
	"github.com/knands42/lorecrafter/internal/utils"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"strings"
	"time"
)

type UserCreationInput struct {
	Username string `json:"username" example:"johndoe"`
	Email    string `json:"email" example:"john@mail.com"`
	Password string `json:"password" example:"12345678"`
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

type User struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	AvatarUrl   string    `json:"avatar_url"`
	LastLoginAt time.Time `json:"last_login_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromSqlcUserToDomain(userSqlc sqlc.User) User {
	return User{
		ID:          userSqlc.ID.Bytes,
		Username:    userSqlc.Username,
		Email:       userSqlc.Email,
		AvatarUrl:   userSqlc.AvatarUrl.String,
		LastLoginAt: userSqlc.LastLoginAt.Time,
		CreatedAt:   userSqlc.CreatedAt.Time,
		UpdatedAt:   userSqlc.UpdatedAt.Time,
	}
}
