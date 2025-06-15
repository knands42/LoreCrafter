package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
)

var (
	ErrInvalidUUID     = fmt.Errorf("invalid UUID")
	UUIDCreationFailed = fmt.Errorf("failed to create UUID")
)

func GeneratePGUUID() (pgtype.UUID, error) {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		log.Printf("invalid UUID")
		return pgtype.UUID{}, UUIDCreationFailed
	}

	return pgtype.UUID{
		Bytes: uuidV7,
		Valid: true,
	}, nil
}

func GeneratePGUUIDFromCustomId(customUUID uuid.UUID) (pgtype.UUID, error) {
	return pgtype.UUID{
		Bytes: customUUID,
		Valid: true,
	}, nil
}

func FromPGTypeUUID(pgu pgtype.UUID) (uuid.UUID, error) {
	if !pgu.Valid {
		log.Printf("invalid UUID %v", pgu)
		return uuid.Nil, ErrInvalidUUID
	}
	return pgu.Bytes, nil
}
