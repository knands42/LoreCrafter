package interfaces

import (
	"context"
)

// Worker defines the interface for background job processing
type Worker interface {
	// Publish publishes a message to a channel
	Publish(ctx context.Context, channel string, message string) error

	// Subscribe subscribes to a channel and executes the handler function when a message is received
	Subscribe(ctx context.Context, channel string, handler func(message string) error) error

	// Close closes the worker connection
	Close() error
}