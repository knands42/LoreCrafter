package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
)

var (
	ErrInvalidUUID = fmt.Errorf("invalid UUID")
)

func FromPGTypeUUID(pgu pgtype.UUID) (uuid.UUID, error) {
	if !pgu.Valid {
		log.Printf("invalid UUID %v", pgu)
		return uuid.Nil, ErrInvalidUUID
	}
	return pgu.Bytes, nil
}
