package main

import (
	"github.com/knands42/lorecrafter/app/api"
	"github.com/knands42/lorecrafter/internal/adapter/database"
	"github.com/knands42/lorecrafter/internal/adapter/database/migrations"
	"github.com/knands42/lorecrafter/internal/adapter/llms"
	"github.com/knands42/lorecrafter/internal/config"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"github.com/tmc/langchaingo/llms/openai"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set up the database
	pgConn, err := database.NewPostgresConnection(&cfg)
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get cur directory: %v", err)
	}

	migrationPath := filepath.Join(cwd, "./internal/adapter/database/migrations")
	migrationsDir := os.DirFS(migrationPath)
	migrations.Up(cfg, migrationsDir)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	repo := sqlc.New(pgConn)

	// Set up the LLM models
	llm, err := openai.New(
		openai.WithToken(cfg.OpenAIAPIKey),
		openai.WithModel("gpt-4-turbo-preview"),
	)
	if err != nil {
		log.Fatalf("failed to create OpenAI client: %v", err)
	}
	llmFactory := llms.NewLlmFactory(llm)

	// Set up the HTTP server
	api.NewServer(cfg, repo, llmFactory).Start()
}
