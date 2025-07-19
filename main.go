package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/knands42/lorecrafter/app/api"
	"github.com/knands42/lorecrafter/internal/adapter/database"
	"github.com/knands42/lorecrafter/internal/adapter/database/migrations"
	"github.com/knands42/lorecrafter/internal/adapter/email"
	"github.com/knands42/lorecrafter/internal/adapter/llms"
	"github.com/knands42/lorecrafter/internal/adapter/security"
	"github.com/knands42/lorecrafter/internal/config"
	"github.com/knands42/lorecrafter/internal/usecases"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
)

func main() {
	// load configuration
	ctx := context.Background()
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set up the database
	pgConn, err := database.NewPostgresConnection(&cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get cur directory: %v", err)
	}

	migrationPath := filepath.Join(cwd, "./internal/adapter/database/migrations")
	migrationsDir := os.DirFS(migrationPath)
	migrations.Up(cfg, migrationsDir)

	repo := sqlc.New(pgConn)

	// Set up the LLM models
	llmFactory, err := llms.NewLlmFactory(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to initialize llms: %v", err)
	}

	// setup adapters
	//worker := worker2.NewValkeyAdapter(cfg)
	//worker.Close()
	tokenMakerAdapter, err := security.NewTokenMakerAdapter(cfg.PrivateKey, cfg.PublicKey)
	if err != nil {
		log.Fatalf("Failed to create token maker: %v", err)
	}
	argon2Adapter := security.NewArgon2Adapter(cfg.PasswordSalt)
	emailSender := email.NewEmailSenderAdapter(cfg.EmailAPIKEY, cfg.EmailDomain)
	templateManager, err := email.NewTemplateManagerAdapter()
	if err != nil {
		log.Fatalf("Failed to create template manager: %v", err)
	}

	// setup usecases
	emailUseCase := usecases.NewEmailUseCase(ctx, emailSender, templateManager, repo, "")
	authUseCase := usecases.NewAuthUseCase(ctx, repo, tokenMakerAdapter, argon2Adapter, cfg.TokenExpiry, emailUseCase)
	aiCampaignUseCase := usecases.NewAICampaignUseCase(ctx, repo, llmFactory)
	campaignUseCase := usecases.NewCampaignUseCase(ctx, repo, aiCampaignUseCase)
	passwordResetUseCase := usecases.NewPasswordResetUseCase(ctx, repo, emailUseCase, templateManager, argon2Adapter, cfg.TokenExpiry)
	campaignInvitationUseCase := usecases.NewCampaignInvitationUseCase(ctx, repo)
	campaignMembersUseCase := usecases.NewCampaignMembersUseCase(ctx, repo)

	// Set up the HTTP server
	server := api.NewServer(
		cfg,
		repo,
		authUseCase,
		campaignUseCase,
		passwordResetUseCase,
		campaignInvitationUseCase,
		campaignMembersUseCase,
	)

	// Start the server with or without TLS
	server.Start()
	//defer pgConn.Close()
}
