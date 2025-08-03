package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/knands42/lorecrafter/internal/config"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"log"
)

// NotificationPayload represents the structure of the notification payload from PostgreSQL
type NotificationPayload struct {
	ID        string                  `json:"id"`
	UserID    string                  `json:"user_id"`
	Type      sqlc.NotificationType   `json:"type"`
	Status    sqlc.NotificationStatus `json:"status"`
	Payload   []byte                  `json:"payload"`
	CreatedAt string                  `json:"created_at"`
}

// Worker handles Postgresql notification listening and processing
type Worker struct {
	ctx     context.Context
	config  *config.Config
	pool    *pgxpool.Pool
	repo    sqlc.Querier
	helpers Helpers
}

// NewWorker creates a new worker instance
func NewWorker(ctx context.Context, pool *pgxpool.Pool, repo sqlc.Querier, cfg *config.Config) *Worker {
	return &Worker{
		ctx:     ctx,
		config:  cfg,
		pool:    pool,
		repo:    repo,
		helpers: Helpers{},
	}
}

// Start begins listening for PostgreSQL notifications
func (w *Worker) Start() error {
	conn, err := w.pool.Acquire(w.ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()

	// Listen for notifications on the 'new_notification' channel
	_, err = conn.Exec(w.ctx, "LISTEN new_notification")
	if err != nil {
		return fmt.Errorf("failed to listen on channel: %w", err)
	}

	log.Println("Worker started listening for notifications on 'new_notification' channel")

	// Process notifications in a loop
	for {
		notification, err := conn.Conn().WaitForNotification(w.ctx)
		if err != nil {
			log.Printf("Error waiting for notification: %v", err)
			continue
		}

		log.Printf("Received notification on channel %s: %s", notification.Channel, notification.Payload)

		// Process the notification
		if err := w.processNotification(notification.Payload); err != nil {
			log.Printf("Error processing notification: %v", err)
		}
	}
}

// processNotification handles the notification payload and routes it to the appropriate handler
func (w *Worker) processNotification(payload string) error {
	var notificationPayload NotificationPayload
	if err := json.Unmarshal([]byte(payload), &notificationPayload); err != nil {
		return fmt.Errorf("failed to unmarshal notification payload: %w", err)
	}

	log.Printf("Processing notification of type: %s", notificationPayload.Type)

	// Handle different notification types
	switch notificationPayload.Type {
	case sqlc.NotificationTypeCampaignInvite:
		return w.handleCampaignInvite(notificationPayload)
	default:
		log.Printf("Unhandled notification type: %s", notificationPayload.Type)
		return nil
	}
}
