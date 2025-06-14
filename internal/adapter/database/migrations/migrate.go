package migrations

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/knands42/lorecrafter/internal/config"
	"log"
	"maragu.dev/migrate"
	"os"
)

var migrationsDir = os.DirFS(".")

func Up(cfg config.Config) {
	db, _ := sql.Open("pgx", cfg.PostgresURL)
	if err := migrate.Up(context.Background(), db, migrationsDir); err != nil {
		panic(err)
	}

	log.Printf("Successfully migrated database")
}

func Down(cfg config.Config) {
	db, _ := sql.Open("pgx", cfg.PostgresURL)
	if err := migrate.Down(context.Background(), db, migrationsDir); err != nil {
		panic(err)
	}
}
