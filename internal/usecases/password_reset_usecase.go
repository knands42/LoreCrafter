package usecases

import (
	"context"
	"errors"
	"fmt"
	"github.com/knands42/lorecrafter/internal/utils"
	"log"
	"time"

	"github.com/knands42/lorecrafter/internal/domain"
	"github.com/knands42/lorecrafter/internal/interfaces"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
)

var (
	ErrFailedToGenerateToken            = errors.New("failed to generate token")
	ErrFailedToCreatePasswordResetToken = errors.New("failed to create password reset token")
	ErrInvalidResetToken                = errors.New("invalid or expired password reset token")
	ErrPasswordResetTokenExpired        = errors.New("password reset token has expired")
	ErrPasswordResetTokenUsed           = errors.New("password reset token has already been used")
)

type PasswordResetUseCase struct {
	ctx          context.Context
	repo         sqlc.Querier
	emailUseCase *EmailUseCase
	tm           interfaces.TemplateManager
	argon2Hash   interfaces.Argon2Hash
	tokenExpiry  time.Duration
}

func NewPasswordResetUseCase(
	ctx context.Context,
	repo sqlc.Querier,
	emailUseCase *EmailUseCase,
	tm interfaces.TemplateManager,
	argon2Hash interfaces.Argon2Hash,
	tokenExpiry time.Duration,
) *PasswordResetUseCase {
	return &PasswordResetUseCase{
		ctx:          ctx,
		repo:         repo,
		emailUseCase: emailUseCase,
		argon2Hash:   argon2Hash,
		tokenExpiry:  tokenExpiry,
	}
}

// RequestPasswordReset handles the password reset request
// It generates a token, saves it to the database, and sends a password reset email
func (uc *PasswordResetUseCase) RequestPasswordReset(input domain.ForgotPasswordInput) error {
	// Invalidate any existing tokens for this user
	if err := uc.repo.InvalidateAllUserTokens(uc.ctx, input.Email); err != nil {
		log.Printf("Failed to invalidate existing tokens: %v", err)
		log.Printf("Continuing anyway, as this is not a critical error")
	}

	// Create new token
	token, err := utils.GenerateRandomHexString(32)
	if err != nil {
		return ErrFailedToGenerateToken
	}
	tokenInput := domain.NewCreatePasswordResetTokenInput(input.Email, token)
	params, err := tokenInput.PrepareToInsert()
	if err != nil {
		return ErrFailedToCreatePasswordResetToken
	}

	dbToken, err := uc.repo.CreatePasswordResetToken(uc.ctx, params)
	if err != nil {
		return ErrFailedToCreatePasswordResetToken
	}

	// Send password reset email
	go func() {
		if err = uc.emailUseCase.SendPasswordResetTokenEmail(domain.SendEmailVerificationToken{Token: token, Email: input.Email}); err != nil {
			log.Printf("Failed to send password reset email: %v", err)
		}
	}()

	log.Printf("Password reset token created for user %s (token ID: %v)", input.Email, dbToken.ID)
	return nil
}

// ResetPassword handles the password reset process
func (uc *PasswordResetUseCase) ResetPassword(input domain.PasswordResetInput) error {
	// Update the user's password
	hashedPassword, err := uc.argon2Hash.HashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	_, err = uc.repo.UpdateUserPasswordFromToken(uc.ctx, sqlc.UpdateUserPasswordFromTokenParams{
		Token:          input.Token,
		HashedPassword: hashedPassword,
	})
	if err != nil && err.Error() == "no rows in result set" {
		return ErrInvalidResetToken
	}

	// Mark the token as used
	go func() {
		if err := uc.repo.InvalidatePasswordResetToken(uc.ctx, input.Token); err != nil {
			log.Printf("Failed to invalidate password reset token: %v", err)
		}
	}()

	log.Printf("Password reset successful for user ID: %s", input.Email)
	return nil
}
