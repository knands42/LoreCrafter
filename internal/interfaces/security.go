package interfaces

import (
	"github.com/knands42/lorecrafter/internal/domain"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"time"
)

type TokenMaker interface {
	CreateToken(user sqlc.User, duration time.Duration) (string, time.Time, error)
	VerifyToken(token string) (*domain.TokenPayload, error)
}

type Argon2Hash interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password, encodedHash string) (bool, error)
}
