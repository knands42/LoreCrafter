package workers

import (
	"fmt"
	"github.com/knands42/lorecrafter/internal/domain"
	"log"
)

// handleCampaignInvite processes campaign invitation notifications
func (w *Worker) handleCampaignInvite(payload NotificationPayload) error {
	log.Printf("Handling campaign invite notification for user: %s", payload.UserID)

	// Parse the UUID strings
	userID, err := w.helpers.parseUUID(payload.UserID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	notificationID, err := w.helpers.parseUUID(payload.ID)
	if err != nil {
		return fmt.Errorf("invalid notification ID: %w", err)
	}

	// Convert the notification to a domain model
	notification := domain.Notification{
		ID:        notificationID,
		UserID:    userID,
		Type:      payload.Type,
		Status:    payload.Status,
		Payload:   string(payload.Payload),
		CreatedAt: w.helpers.parseTime(payload.CreatedAt),
	}

	// TODO: send an email, or preload into a in-memory cache
	log.Printf("Processed campaign invite notification: %+v", notification)

	return nil
}
