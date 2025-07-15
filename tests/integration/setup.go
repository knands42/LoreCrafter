package integration

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
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

var TestDB *pgxpool.Pool
var TestServer *httptest.Server
var TestClient *http.Client

// SetupIntegrationTest sets up the integration test environment
func SetupIntegrationTest() error {
	ctx := context.Background()

	cfg, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	cfg.PostgresURL = "postgres://postgres:postgres@localhost:5433/lorecrafter_test?sslmode=disable"

	// Set up the database
	pgConn, err := database.NewPostgresConnection(&cfg)
	cwd, _ := os.Getwd()
	migrationPath := filepath.Join(cwd, "../../internal/adapter/database/migrations")
	migrationsDir := os.DirFS(migrationPath)
	migrations.Up(cfg, migrationsDir)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	repo := sqlc.New(pgConn)

	// Set up the LLM models
	llm, err := llms.GetGeminiLlmFactory(ctx, cfg)
	
	if err != nil {
		log.Fatalf("failed to create OpenAI client: %v", err)
	}
	llmFactory := llms.NewLlmFactory(llm)

	// setup adapters
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

	server := api.NewServer(
		cfg,
		repo,
		authUseCase,
		campaignUseCase,
		passwordResetUseCase,
	)

	TestDB = pgConn
	TestServer = httptest.NewServer(server.Router)
	TestClient = TestServer.Client()

	return nil
}

// TeardownIntegrationTest tears down the integration test environment
func TeardownIntegrationTest() {
	if TestServer != nil {
		TestServer.Close()
	}

	if TestDB != nil {
		TestDB.Close()
	}
}
