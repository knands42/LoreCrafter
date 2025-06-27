package usecases

import (
	"context"
	"errors"
	"fmt"

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
	if err := eu.sender.SendEmail(input.Email, subject, body); err != nil {
		return ErrSendEmail
	}

	return nil
}

func (eu *EmailUseCase) ValidateEmailVerificationToken(token string) error {
	return nil
}
