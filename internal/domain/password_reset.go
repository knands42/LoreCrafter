package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/knands42/lorecrafter/internal/utils"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
)

type ForgotPasswordInput struct {
	Email string `json:"email" example:"johndoe@mail.com"`
}

type CreatePasswordResetTokenInput struct {
	Email     string `json:"email" example:"johndoe@mail.com"`
	TokenHash string `json:"token_hash"`
	expiresAt time.Time
}

func NewCreatePasswordResetTokenInput(email, token string) *CreatePasswordResetTokenInput {
	return &CreatePasswordResetTokenInput{
		Email:     email,
		TokenHash: token,
		expiresAt: time.Now().Add(24 * time.Hour), // Token expires in 24 hours
	}
}

// PrepareToInsert prepares the input for database insertion
func (input *CreatePasswordResetTokenInput) PrepareToInsert() (sqlc.CreatePasswordResetTokenParams, error) {
	newUUID, err := utils.GeneratePGUUID()
	if err != nil {
		return sqlc.CreatePasswordResetTokenParams{}, err
	}

	return sqlc.CreatePasswordResetTokenParams{
		ID:        newUUID,
		Email:     input.Email,
		TokenHash: input.TokenHash,
		ExpiresAt: pgtype.Timestamptz{Time: input.expiresAt, Valid: true},
	}, nil
}

// PasswordResetToken represents a password reset token in the domain
// This is used to convert between database and domain models
type PasswordResetToken struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	TokenHash string    `json:"-"`
	ExpiresAt time.Time `json:"expires_at"`
	Used      bool      `json:"used"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// FromSqlcPasswordResetToken converts a database model to a domain model
func FromSqlcPasswordResetToken(token sqlc.PasswordResetToken) PasswordResetToken {
	return PasswordResetToken{
		ID:        token.ID.Bytes,
		UserID:    token.UserID.Bytes,
		TokenHash: token.TokenHash,
		ExpiresAt: token.ExpiresAt.Time,
		Used:      token.Used,
		CreatedAt: token.CreatedAt.Time,
		UpdatedAt: token.UpdatedAt.Time,
	}
}

type PasswordResetInput struct {
	Token    string `json:"token"`
	Email    string `json:"email" example:"johndoe@mail.com"`
	Password string `json:"password" example:"87654321"`
}

func NewPasswordResetInput(token, email, password string) *PasswordResetInput {
	return &PasswordResetInput{
		Token:    token,
		Email:    email,
		Password: password,
	}
}
