package usecases

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"log"
	"time"

	"github.com/knands42/lorecrafter/internal/adapter/security"
	"github.com/knands42/lorecrafter/internal/domain"
)

var (
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrCheckingPassword     = errors.New("error checking password")
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrCreateUser           = errors.New("error creating user")
	ErrCheckingIfUserExists = errors.New("error checking if user exists")
	ErrHashPassword         = errors.New("error hashing password")
)

type AuthUseCase struct {
	ctx         context.Context
	userRepo    sqlc.Querier
	tokenMaker  *security.TokenMaker
	tokenExpiry time.Duration
}

func NewAuthUseCase(
	ctx context.Context,
	userRepo sqlc.Querier,
	tokenMaker *security.TokenMaker,
	tokenExpiry time.Duration,
) *AuthUseCase {
	return &AuthUseCase{
		ctx:         ctx,
		userRepo:    userRepo,
		tokenMaker:  tokenMaker,
		tokenExpiry: tokenExpiry,
	}
}

// Register creates a new user and generates a token for them
func (uc *AuthUseCase) Register(input domain.UserCreationInput) (*domain.AuthOutput, error) {
	validationErrors := input.Validate()
	if validationErrors != nil {
		log.Printf("validation errors: %v", validationErrors)
		return nil, validationErrors
	}

	getUserByUsernameOrEmailParams := sqlc.GetUserByUsernameOrEmailParams{
		Username: input.Username,
		Email:    input.Email,
	}
	_, err := uc.userRepo.GetUserByUsernameOrEmail(uc.ctx, getUserByUsernameOrEmailParams)
	if err == nil {
		return nil, ErrUserAlreadyExists
	}

	hashedPassword, err := security.HashPassword(input.Password, nil)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, ErrHashPassword
	}

	newUser, err := domain.NewUser(input.Username, input.Email, hashedPassword)
	if err != nil {
		return nil, err
	}

	// Save the user
	createUserParams := sqlc.CreateUserParams{
		ID: pgtype.UUID{
			Bytes: newUser.ID,
			Valid: true,
		},
		Username:       newUser.Username,
		Email:          newUser.Email,
		HashedPassword: hashedPassword,
	}
	if _, err := uc.userRepo.CreateUser(uc.ctx, createUserParams); err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, ErrCreateUser
	}

	// Generate a token
	token, expiresAt, err := uc.tokenMaker.CreateToken(newUser, uc.tokenExpiry)
	if err != nil {
		return nil, fmt.Errorf("error generating token: %w", err)
	}

	return &domain.AuthOutput{
		User:      *newUser,
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}

// Login authenticates a user and generates a token for them
func (uc *AuthUseCase) Login(req domain.LoginInput) (*domain.AuthOutput, error) {
	// Get the user by username
	dbUser, err := uc.userRepo.GetUserByUsername(uc.ctx, req.Username)
	if err != nil {
		log.Printf("error getting user: %v", err)
		return nil, ErrCheckingIfUserExists
	}

	// Verify the password
	match, err := security.VerifyPassword(req.Password, dbUser.HashedPassword)
	if err != nil {
		log.Printf("error verifying password: %v", err)
		return nil, ErrCheckingPassword
	}

	if !match {
		return nil, ErrInvalidCredentials
	}

	// Generate a token
	userDomain, err := domain.ToDomainUser(dbUser)
	if err != nil {
		return nil, err
	}
	token, expiresAt, err := uc.tokenMaker.CreateToken(&userDomain, uc.tokenExpiry)
	if err != nil {
		return nil, fmt.Errorf("error generating token: %w", err)
	}

	return &domain.AuthOutput{
		User:      userDomain,
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}

// VerifyToken verifies a token and returns the payload
func (uc *AuthUseCase) VerifyToken(token string) (*domain.TokenPayload, error) {
	return uc.tokenMaker.VerifyToken(token)
}
