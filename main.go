package main

import (
	"context"
	"github.com/knands42/lorecrafter/cmd/api"
	"github.com/knands42/lorecrafter/cmd/api/routes"
	"github.com/knands42/lorecrafter/internal/adapter/database"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/knands42/lorecrafter/internal/adapter/security"
	"github.com/knands42/lorecrafter/internal/usecases"
)

// Configuration constants
const (
	ServerPort     = "8080"
	TokenExpiry    = 24 * time.Hour
	PrivateKeyPath = "private_key.pem"
	PublicKeyPath  = "public_key.pem"
)

func main() {
	// Set up repository
	// TODO: Get from env using viper
	pgConn, err := database.NewPostgresConnection()
	repo := sqlc.New(pgConn)

	// Set up token maker
	tokenMaker, err := security.NewTokenMaker(PrivateKeyPath, PublicKeyPath)
	if err != nil {
		log.Fatalf("Failed to create token maker: %v", err)
	}

	// Set up use cases
	ctx := context.Background()
	authUseCase := usecases.NewAuthUseCase(ctx, repo, tokenMaker, TokenExpiry)

	// Set up HTTP handlers
	authHandler := routes.NewAuthHandler(authUseCase)

	// Set up HTTP server
	server := api.NewServer(ServerPort)
	api.SetupRoutes(server, authHandler)

	// Start the server in a goroutine
	go func() {
		log.Printf("Starting server on port %s", ServerPort)
		if err := server.Start(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
