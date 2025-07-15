package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/knands42/lorecrafter/internal/config"
	"github.com/knands42/lorecrafter/internal/interfaces"
	"github.com/valkey-io/valkey-go"
)

// ValkeyAdapter implements the Worker interface using Valkey
type ValkeyAdapter struct {
	client valkey.Client
}

// NewValkeyAdapter creates a new ValkeyAdapter with the given configuration
func NewValkeyAdapter(cfg config.Config) interfaces.Worker {
	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{cfg.VALKEY_ADDRESS}})
	if err != nil {
		panic(err)
	}

	return &ValkeyAdapter{
		client: client,
	}
}

// Publish publishes a message to a channel
func (v *ValkeyAdapter) Publish(ctx context.Context, channel string, message string) error {
	fmt.Printf("Publishing message to channel %s: %s\n", channel, message)

	cmd := v.client.B().Publish().Channel(channel).Message(message).Build()
	err := v.client.Do(ctx, cmd).Error()
	if err != nil {
		return fmt.Errorf("failed to publish message to channel %s: %w", channel, err)
	}

	return nil
}

// Subscribe subscribes to a channel and executes the handler function when a message is received
func (v *ValkeyAdapter) Subscribe(ctx context.Context, channel string, handler func(message string) error) error {
	fmt.Printf("Subscribing to channel: %s\n", channel)

	cmd := v.client.B().Subscribe().Channel(channel).Build()
	err := v.client.Do(ctx, cmd).Error()
	if err != nil {
		return fmt.Errorf("failed to subscribe to channel %s: %w", channel, err)
	}

	// Start a goroutine to listen for messages
	go func() {
		for {
			select {
			case <-ctx.Done():
				unsubCmd := v.client.B().Unsubscribe().Channel(channel).Build()
				_ = v.client.Do(ctx, unsubCmd).Error()
				return
			default:
				// Poll for messages
				getCmd := v.client.B().Get().Key(channel).Build()
				result, err := v.client.Do(ctx, getCmd).ToString()
				if err == nil && result != "" {
					// Process the message
					if err := handler(result); err != nil {
						fmt.Printf("Error handling message: %v\n", err)
					}
				}
				// Sleep briefly to avoid tight polling
				time.Sleep(time.Second)
			}
		}
	}()

	<-ctx.Done()
	return ctx.Err()
}

// Close closes the Valkey client connection
func (v *ValkeyAdapter) Close() error {
	fmt.Println("Closing Valkey client")
	v.client.Close()
	return nil
}
