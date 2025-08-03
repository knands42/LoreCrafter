package workers

import (
	"github.com/google/uuid"
	"log"
	"time"
)

type Helpers struct{}

// Helper function to parse UUID strings
func (h *Helpers) parseUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}

// Helper function to parse time strings
func (h *Helpers) parseTime(timeStr string) time.Time {
	t, err := time.Parse(time.RFC3339Nano, timeStr)
	if err == nil {
		return t
	}

	t, err = time.Parse(time.RFC3339, timeStr)
	if err == nil {
		return t
	}

	t, err = time.Parse("2006-01-02T15:04:05.999999", timeStr)
	if err != nil {
		log.Printf("Error parsing time: %v", err)
		return time.Now()
	}
	return t
}
