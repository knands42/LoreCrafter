package usecases

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/knands42/lorecrafter/internal/domain"
	"github.com/knands42/lorecrafter/internal/interfaces"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
)

var (
	ErrSendVerificationEmail = errors.New("failed to send verification email")
	ErrEmailTemplateRender   = errors.New("failed to render email template")
	ErrSendEmail             = errors.New("failed to send email")
)

type EmailUseCase struct {
	ctx     context.Context
	sender  interfaces.EmailSender
	tm      interfaces.TemplateManager
	baseURL string
	repo    sqlc.Querier
}

func NewEmailUseCase(ctx context.Context, sender interfaces.EmailSender, tm interfaces.TemplateManager, repo sqlc.Querier, baseURL string) *EmailUseCase {
	return &EmailUseCase{
		ctx:     ctx,
		sender:  sender,
		tm:      tm,
		baseURL: baseURL,
		repo:    repo,
	}
}

func (eu *EmailUseCase) CreateEmailVerificationToken(input domain.CreateEmailVerificationTokenInput) (string, error) {
	params, err := input.PrepareToInsert()
	if err != nil {
		return "", err
	}
	result, err := eu.repo.CreateEmailVerificationToken(eu.ctx, params)
	if err != nil {
		return "", err
	}

	return result.EmailVerificationToken.String, nil
}

func (eu *EmailUseCase) SendVerificationEmail(input domain.SendEmailVerificationToken) error {
	verificationURL := fmt.Sprintf("%s/api/v1/auth/verify-email?token=%s&email=%s", eu.baseURL, input.Token, input.Email)

	body, err := eu.tm.Render("email_verification", map[string]string{
		"VerificationURL": verificationURL,
	})
	if err != nil {
		return ErrEmailTemplateRender
	}

	subject := "Verify Your Email Address"
	if err := eu.sender.SendEmail(input.Name, input.Email, subject, body); err != nil {
		log.Printf("Failed to send verification email: %v", err)
		return ErrSendEmail
	}

	return nil
}

// TODO: implement
func (eu *EmailUseCase) ValidateEmailVerificationToken(token string) error {
	return nil
}

func (eu *EmailUseCase) SendPasswordResetTokenEmail(input domain.SendEmailVerificationToken) error {
	body, err := eu.tm.Render("password_reset", map[string]string{
		"ResetURL": fmt.Sprintf("%s/reset-password?token=%s&email=%s", eu.baseURL, input.Token, input.Email),
	})
	if err != nil {
		log.Printf("Failed to render password reset template: %v", err)
		return ErrEmailTemplateRender
	}
	if err := eu.sender.SendEmail(input.Name, input.Email, "Reset Your Password", body); err != nil {
		log.Printf("Failed to send password reset email: %v", err)
		return ErrSendEmail
	}

	return nil
}
