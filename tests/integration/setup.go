package integration

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/knands42/lorecrafter/app/api"
	"github.com/knands42/lorecrafter/internal/adapter/database"
	"github.com/knands42/lorecrafter/internal/adapter/database/migrations"
	"github.com/knands42/lorecrafter/internal/config"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
)

var TestDB *pgxpool.Pool
var TestServer *httptest.Server
var TestClient *http.Client

// SetupIntegrationTest sets up the integration test environment
func SetupIntegrationTest() error {
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

	// Set up the HTTP server
	server := api.NewServer(cfg, repo)

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
