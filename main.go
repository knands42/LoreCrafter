package main

import (
	"encoding/base64"
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/knands42/lorecrafter/app/api"
	"github.com/knands42/lorecrafter/internal/adapter/database"
	"github.com/knands42/lorecrafter/internal/adapter/database/migrations"
	"github.com/knands42/lorecrafter/internal/adapter/llms"
	"github.com/knands42/lorecrafter/internal/config"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	// Parse command line flags
	enableTLS := flag.Bool("tls", false, "Enable TLS/HTTPS")
	flag.Parse()

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
	server := api.NewServer(cfg, repo, llmFactory)

	// Start the server with or without TLS
	if *enableTLS {
		certFile, keyFile := prepareCertificates(cfg.SSLCert, cfg.SSLKey)
		server.StartTLS(certFile, keyFile)
	} else {
		server.Start()
	}
}

func prepareCertificates(certFile, keyFile string) (string, string) {
	// decode base64
	certFileDecoded, err := base64.StdEncoding.DecodeString(certFile)
	if err != nil {
		log.Fatalf("Failed to decode SSL cert: %v", err)
	}
	keyFileDecoded, err := base64.StdEncoding.DecodeString(keyFile)
	if err != nil {
		log.Fatalf("Failed to decode SSL key: %v", err)
	}

	// get cwd and append with certs
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get cur directory: %v", err)
	}

	err = os.WriteFile(filepath.Join(cwd, "certs", "server.crt"), certFileDecoded, 0644)
	if err != nil {
		log.Fatalf("Failed to write SSL cert: %v", err)
	}
	err = os.WriteFile(filepath.Join(cwd, "certs", "server.key"), keyFileDecoded, 0644)
	if err != nil {
		log.Fatalf("Failed to write SSL key: %v", err)
	}

	certPath := filepath.Join(cwd, "certs", "server.crt")
	keyPath := filepath.Join(cwd, "certs", "server.key")
	return certPath, keyPath
}
