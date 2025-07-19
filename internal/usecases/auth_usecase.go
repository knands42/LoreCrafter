package usecases

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/knands42/lorecrafter/internal/interfaces"
	"github.com/knands42/lorecrafter/internal/utils"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"

	"github.com/knands42/lorecrafter/internal/domain"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrCheckingPassword   = errors.New("error checking password")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrCreateUser         = errors.New("error creating user")
	ErrUserNotFound       = errors.New("user not found")
	ErrHashPassword       = errors.New("error hashing password")
)

type AuthUseCase struct {
	ctx          context.Context
	userRepo     sqlc.Querier
	tokenMaker   interfaces.TokenMaker
	argon2Hash   interfaces.Argon2Hash
	tokenExpiry  time.Duration
	emailUseCase *EmailUseCase
}

func NewAuthUseCase(
	ctx context.Context,
	userRepo sqlc.Querier,
	tokenMaker interfaces.TokenMaker,
	argon2Hash interfaces.Argon2Hash,
	tokenExpiry time.Duration,
	emailUseCase *EmailUseCase,
) *AuthUseCase {
	return &AuthUseCase{
		ctx:          ctx,
		userRepo:     userRepo,
		tokenMaker:   tokenMaker,
		argon2Hash:   argon2Hash,
		tokenExpiry:  tokenExpiry,
		emailUseCase: emailUseCase,
	}
}

// Register creates a new user and generates a token for them
func (uc *AuthUseCase) Register(input domain.UserCreationInput) (domain.User, error) {
	validationErrors := input.Validate()
	if validationErrors != nil {
		log.Printf("validation errors: %v", validationErrors)
		return domain.User{}, validationErrors
	}

	getUserByUsernameOrEmailParams := sqlc.GetUserByUsernameOrEmailParams{
		Username: input.Username,
		Email:    input.Email,
	}
	_, err := uc.userRepo.GetUserByUsernameOrEmail(uc.ctx, getUserByUsernameOrEmailParams)
	if err == nil {
		return domain.User{}, ErrUserAlreadyExists
	}

	hashedPassword, err := uc.argon2Hash.HashPassword(input.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return domain.User{}, ErrHashPassword
	}

	// Save the user
	createUserParams, err := input.ToSqlcParams(hashedPassword)
	if err != nil {
		return domain.User{}, err
	}

	createdUser, err := uc.userRepo.CreateUser(uc.ctx, createUserParams)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return domain.User{}, ErrCreateUser
	}

	// Send verification email
	createdUserUUID, err := utils.FromPGTypeUUID(createdUser.ID)
	if err != nil {
		log.Printf("Error converting new user ID to UUID: %v", err)
	}
	createdEmailVerificationTokenInput := domain.NewCreateEmailVerificationTokenInput(createdUserUUID)
	emailVerificationToken, err := uc.emailUseCase.CreateEmailVerificationToken(*createdEmailVerificationTokenInput)
	if err != nil {
		log.Printf("Error creating email verification token: %v", err)
	}

	go func() {
		if err := uc.emailUseCase.SendVerificationEmail(*domain.NewSendEmailVerificationToken(createdUser.Username, emailVerificationToken, createdUser.Email)); err != nil {
			log.Printf("Error sending verification email: %v", err)
		}
	}()

	return domain.FromSqlcUserToDomain(createdUser), nil
}

// Login authenticates a user and generates a token for them
func (uc *AuthUseCase) Login(req domain.LoginInput) (domain.AuthOutput, error) {
	// Validate the input
	if err := req.Validate(); err != nil {
		return domain.AuthOutput{}, err
	}

	// Get the user by username
	user, err := uc.userRepo.GetUserByUsernameOrEmail(uc.ctx, sqlc.GetUserByUsernameOrEmailParams{
		Username: req.UsernameOrEmail,
		Email:    req.UsernameOrEmail,
	})
	if err != nil {
		log.Printf("error getting user: %v", err)
		return domain.AuthOutput{}, ErrUserNotFound
	}

	// Verify the password
	match, err := uc.argon2Hash.VerifyPassword(req.Password, user.HashedPassword)
	if err != nil {
		log.Printf("error verifying password: %v", err)
		return domain.AuthOutput{}, ErrCheckingPassword
	}

	if !match {
		return domain.AuthOutput{}, ErrInvalidCredentials
	}

	// Generate a token
	token, expiresAt, err := uc.tokenMaker.CreateToken(user, uc.tokenExpiry)
	if err != nil {
		return domain.AuthOutput{}, fmt.Errorf("error generating token: %w", err)
	}

	return domain.AuthOutput{
		User:      domain.FromSqlcUserToDomain(user),
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}

// VerifyToken verifies a token and returns the payload
func (uc *AuthUseCase) VerifyToken(token string) (*domain.TokenPayload, error) {
	return uc.tokenMaker.VerifyToken(token)
}

// TODO: implement
func (uc *AuthUseCase) RefreshToken(token string) (domain.TokenPayload, error) {
	return domain.TokenPayload{}, nil
}
