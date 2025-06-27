package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/knands42/lorecrafter/internal/utils"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
)

type CreateEmailVerificationTokenInput struct {
	Token  string    `json:"token"`
	UserId uuid.UUID `json:"user_id"`
}

func NewCreateEmailVerificationTokenInput(userId uuid.UUID) *CreateEmailVerificationTokenInput {
	return &CreateEmailVerificationTokenInput{
		UserId: userId,
	}
}

func (input *CreateEmailVerificationTokenInput) PrepareToInsert() (sqlc.CreateEmailVerificationTokenParams, error) {
	newUUUIDV7, err := utils.GeneratePGUUID()
	if err != nil {
		return sqlc.CreateEmailVerificationTokenParams{}, err
	}
	userUUIDV7, err := utils.GeneratePGUUIDFromCustomId(input.UserId)
	if err != nil {
		return sqlc.CreateEmailVerificationTokenParams{}, err
	}

	if input.Token == "" {
		input.Token, err = utils.GenerateRandomHexString(32)
		if err != nil {
			return sqlc.CreateEmailVerificationTokenParams{}, err
		}
	}

	return sqlc.CreateEmailVerificationTokenParams{
		ID:                         newUUUIDV7,
		UserID:                     userUUIDV7,
		EmailVerificationToken:     pgtype.Text{String: input.Token, Valid: true},
		EmailVerificationSentAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
		EmailVerificationExpiresAt: pgtype.Timestamptz{Time: time.Now().Add(24 * time.Hour), Valid: true},
		CreatedAt:                  pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:                  pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}, nil
}

type SendEmailVerificationToken struct {
	Token string `json:"token"`
	Email string `json:"email"`
}

func NewSendEmailVerificationToken(token string, email string) *SendEmailVerificationToken {
	return &SendEmailVerificationToken{
		Token: token,
		Email: email,
	}
}
