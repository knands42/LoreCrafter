package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/knands42/lorecrafter/internal/adapter/worker"
	"github.com/knands42/lorecrafter/internal/config"
	"github.com/knands42/lorecrafter/internal/interfaces"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create a Valkey worker
	valkeyWorker := worker.NewValkeyAdapter(cfg)
	defer valkeyWorker.Close()

	// Create a context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up signal handling for graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Println("Shutting down...")
		cancel()
	}()

	// Start a goroutine to publish messages
	go publishMessages(ctx, valkeyWorker)

	// Start a goroutine to consume messages
	go consumeMessages(ctx, valkeyWorker)

	// Wait for context to be cancelled
	<-ctx.Done()
	fmt.Println("Exiting...")
}

// publishMessages publishes a message to the "example-channel" every second
func publishMessages(ctx context.Context, worker interfaces.Worker) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	count := 0
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			message := fmt.Sprintf("Message %d", count)
			err := worker.Publish(ctx, "example-channel", message)
			if err != nil {
				log.Printf("Failed to publish message: %v", err)
			} else {
				log.Printf("Published message: %s", message)
			}
			count++
		}
	}
}

// consumeMessages subscribes to the "example-channel" and processes messages
func consumeMessages(ctx context.Context, worker interfaces.Worker) {
	err := worker.Subscribe(ctx, "example-channel", func(message string) error {
		log.Printf("Received message: %s", message)
		return nil
	})
	if err != nil && err != context.Canceled {
		log.Printf("Subscription ended with error: %v", err)
	}
}