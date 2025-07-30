package domain

import (
	"time"

	"github.com/google/uuid"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
)

type Notification struct {
	ID        uuid.UUID               `json:"id"`
	UserID    uuid.UUID               `json:"user_id"`
	Type      sqlc.NotificationType   `json:"type"`
	Payload   string                  `json:"payload"`
	Status    sqlc.NotificationStatus `json:"status"`
	CreatedAt time.Time               `json:"created_at"`
	ReadAt    time.Time               `json:"read_at"`
}

func NewNotificationFromSQLC(notification sqlc.Notification) Notification {
	return Notification{
		ID:        notification.ID.Bytes,
		UserID:    notification.UserID.Bytes,
		Type:      notification.Type,
		Payload:   string(notification.Payload),
		Status:    notification.Status,
		CreatedAt: notification.CreatedAt.Time,
		ReadAt:    notification.ReadAt.Time,
	}
}
