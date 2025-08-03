package usecases

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/knands42/lorecrafter/internal/domain"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
)

var (
	ErrGettingNotifications = errors.New("could not get notifications")
)

type NotificationUseCase struct {
	ctx  context.Context
	repo sqlc.Querier
}

func NewNotificationUseCase(
	ctx context.Context,
	repo sqlc.Querier,
) *NotificationUseCase {
	return &NotificationUseCase{
		ctx:  ctx,
		repo: repo,
	}
}

// GetNotifications retrieves notifications for a specific user
func (n *NotificationUseCase) GetNotifications(userID uuid.UUID) ([]domain.Notification, error) {
	pgUUID := pgtype.UUID{
		Bytes: userID,
		Valid: true,
	}

	notifications, err := n.repo.GetNotifications(n.ctx, pgUUID)
	if err != nil {
		log.Printf("error getting notifications: %v", err)
		return nil, ErrGettingNotifications
	}

	var result []domain.Notification
	for _, notification := range notifications {
		result = append(result, domain.NewNotificationFromSQLC(notification))
	}

	return result, nil
}
