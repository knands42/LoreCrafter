package security

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"

	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

// Argon2Params defines the parameters for Argon2 password hashing
type Argon2Adapter struct {
	salt        string
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func NewArgon2Adapter(salt string) *Argon2Adapter {
	return &Argon2Adapter{
		salt:        salt,
		memory:      64 * 1024, // 64MB
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
}

// HashPassword hashes a password using Argon2id
func (arg *Argon2Adapter) HashPassword(password string) (string, error) {
	hash := argon2.IDKey(
		[]byte(password),
		[]byte(arg.salt),
		arg.iterations,
		arg.memory,
		arg.parallelism,
		arg.keyLength,
	)

	return fmt.Sprintf(
		"%s$%s",
		arg.salt,
		base64.RawStdEncoding.EncodeToString(hash),
	), nil
}

// VerifyPassword checks if a password matches a hash
func (arg *Argon2Adapter) VerifyPassword(password, encodedHash string) (bool, error) {
	hashPassword, err := arg.HashPassword(password)
	if err != nil {
		return false, err
	}

	if subtle.ConstantTimeCompare([]byte(hashPassword), []byte(encodedHash)) == 1 {
		return true, nil
	}

	return false, nil
}
