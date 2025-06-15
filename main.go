package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/knands42/lorecrafter/app/api"
	"github.com/knands42/lorecrafter/app/api/routes"
	"github.com/knands42/lorecrafter/internal/adapter/database"
	"github.com/knands42/lorecrafter/internal/config"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"

	"github.com/knands42/lorecrafter/internal/adapter/security"
	"github.com/knands42/lorecrafter/internal/usecases"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set up repository
	pgConn, err := database.NewPostgresConnection(&cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	repo := sqlc.New(pgConn)

	// Set up adapters
	tokenMakerAdapter, err := security.NewTokenMakerAdapter(cfg.PrivateKey, cfg.PublicKey)
	if err != nil {
		log.Fatalf("Failed to create token maker: %v", err)
	}
	argon2Adapter := security.NewArgon2Adapter()

	// Set up use cases
	ctx := context.Background()
	authUseCase := usecases.NewAuthUseCase(ctx, repo, tokenMakerAdapter, argon2Adapter, cfg.TokenExpiry)
	campaignUseCase := usecases.NewCampaignUseCase(ctx, repo)

	// Set up HTTP handlers
	authHandler := routes.NewAuthHandler(authUseCase)
	userHandler := routes.NewUserHandler()
	campaignHandler := routes.NewCampaignHandler(campaignUseCase)

	// Set up HTTP server
	server := api.NewServer(cfg.ServerPort)
	api.SetupRoutes(server, authHandler, userHandler, campaignHandler, authUseCase)

	// Start the server in a goroutine
	go func() {
		log.Printf("Starting server on port %s", cfg.ServerPort)
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
