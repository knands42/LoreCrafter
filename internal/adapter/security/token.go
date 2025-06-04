package security

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/knands42/lorecrafter/internal/domain"
	"github.com/o1egl/paseto"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidKey   = errors.New("invalid key format")
)

type TokenMaker struct {
	paseto     *paseto.V2
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
}

func NewTokenMaker(privateKeyContent, publicKeyContent string) (*TokenMaker, error) {
	privateKey, err := loadPrivateKey(privateKeyContent)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: %w", err)
	}

	publicKey, err := loadPublicKey(publicKeyContent)
	if err != nil {
		return nil, fmt.Errorf("failed to load public key: %w", err)
	}

	maker := &TokenMaker{
		paseto:     paseto.NewV2(),
		privateKey: privateKey,
		publicKey:  publicKey,
	}

	return maker, nil
}

// CreateToken creates a new token for a specific user and duration
func (maker *TokenMaker) CreateToken(user *domain.User, duration time.Duration) (string, time.Time, error) {
	payload := domain.TokenPayload{
		UserID:    user.ID.String(),
		Username:  user.Username,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	token, err := maker.paseto.Sign(maker.privateKey, payload, nil)
	if err != nil {
		return "", time.Time{}, err
	}

	return token, payload.ExpiresAt, nil
}

// VerifyToken checks if the token is valid and returns the payload
func (maker *TokenMaker) VerifyToken(token string) (*domain.TokenPayload, error) {
	payload := &domain.TokenPayload{}

	err := maker.paseto.Verify(token, maker.publicKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	if time.Now().After(payload.ExpiresAt) {
		return nil, ErrExpiredToken
	}

	return payload, nil
}

// ParseUserID parses the user ID from the token payload
func ParseUserID(payload *domain.TokenPayload) (uuid.UUID, error) {
	userID, err := uuid.Parse(payload.UserID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID: %w", err)
	}
	return userID, nil
}

func loadPrivateKey(key string) (ed25519.PrivateKey, error) {
	decode, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(decode)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, ErrInvalidKey
	}

	privateKey, err := parsePrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func loadPublicKey(key string) (ed25519.PublicKey, error) {
	decode, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(decode)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, ErrInvalidKey
	}

	publicKey, err := parsePublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

func parsePrivateKey(der []byte) (ed25519.PrivateKey, error) {
	// TODO: In a production environment, you might want to use x509.ParsePKCS8PrivateKey
	if len(der) < ed25519.PrivateKeySize {
		return nil, ErrInvalidKey
	}

	// Extract the private key from the DER-encoded data
	// This is a simplified approach and might need adjustment based on the actual key format
	privateKey := ed25519.PrivateKey(der[len(der)-ed25519.PrivateKeySize:])

	return privateKey, nil
}

func parsePublicKey(der []byte) (ed25519.PublicKey, error) {
	// TODO: In a production environment, you might want to use x509.ParsePKIXPublicKey
	if len(der) < ed25519.PublicKeySize {
		return nil, ErrInvalidKey
	}

	// Extract the public key from the DER-encoded data
	// This is a simplified approach and might need adjustment based on the actual key format
	publicKey := ed25519.PublicKey(der[len(der)-ed25519.PublicKeySize:])

	return publicKey, nil
}
