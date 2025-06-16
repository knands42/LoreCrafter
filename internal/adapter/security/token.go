package security

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/knands42/lorecrafter/internal/utils"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"time"

	"github.com/knands42/lorecrafter/internal/domain"
	"github.com/o1egl/paseto"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidKey   = errors.New("invalid key format")
)

type TokenMakerAdapter struct {
	paseto     *paseto.V2
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
}

func NewTokenMakerAdapter(privateKeyContent, publicKeyContent string) (*TokenMakerAdapter, error) {
	maker := &TokenMakerAdapter{}

	privateKey, err := maker.loadPrivateKey(privateKeyContent)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: %w", err)
	}

	publicKey, err := maker.loadPublicKey(publicKeyContent)
	if err != nil {
		return nil, fmt.Errorf("failed to load public key: %w", err)
	}

	maker.paseto = paseto.NewV2()
	maker.privateKey = privateKey
	maker.publicKey = publicKey

	return maker, nil
}

// CreateToken creates a new token for a specific user and duration
func (maker *TokenMakerAdapter) CreateToken(user sqlc.User, duration time.Duration) (string, time.Time, error) {
	convertedUUID, err := utils.FromPGTypeUUID(user.ID)
	if err != nil {
		return "", time.Time{}, utils.ErrInvalidUUID
	}
	payload := domain.TokenPayload{
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}
	payload.ID = convertedUUID
	payload.Email = user.Email
	payload.Username = user.Username
	payload.AvatarUrl = user.AvatarUrl.String
	payload.LastLoginAt = user.LastLoginAt.Time
	payload.CreatedAt = user.CreatedAt.Time
	payload.UpdatedAt = user.UpdatedAt.Time

	token, err := maker.paseto.Sign(maker.privateKey, payload, nil)
	if err != nil {
		return "", time.Time{}, err
	}

	return token, payload.ExpiresAt, nil
}

// VerifyToken checks if the token is valid and returns the payload
func (maker *TokenMakerAdapter) VerifyToken(token string) (*domain.TokenPayload, error) {
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

func (maker *TokenMakerAdapter) loadPrivateKey(key string) (ed25519.PrivateKey, error) {
	decode, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(decode)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, ErrInvalidKey
	}

	privateKey, err := maker.parsePrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func (maker *TokenMakerAdapter) loadPublicKey(key string) (ed25519.PublicKey, error) {
	decode, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(decode)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, ErrInvalidKey
	}

	publicKey, err := maker.parsePublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

func (maker *TokenMakerAdapter) parsePrivateKey(der []byte) (ed25519.PrivateKey, error) {
	key, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		return nil, err
	}

	ed25519Key := key.(ed25519.PrivateKey)
	if len(ed25519Key) < ed25519.PrivateKeySize {
		return nil, ErrInvalidKey
	}

	return ed25519Key, nil
}

func (maker *TokenMakerAdapter) parsePublicKey(der []byte) (ed25519.PublicKey, error) {
	key, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		return nil, err
	}

	ed25519Key := key.(ed25519.PublicKey)
	if len(ed25519Key) < ed25519.PublicKeySize {
		return nil, ErrInvalidKey
	}

	return ed25519Key, nil
}
