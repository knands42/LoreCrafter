package migrations

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/knands42/lorecrafter/internal/config"
	"io/fs"
	"log"
	"maragu.dev/migrate"
)

func Up(cfg config.Config, migrationsDir fs.FS) {
	db, _ := sql.Open("pgx", cfg.PostgresURL)
	if err := migrate.Up(context.Background(), db, migrationsDir); err != nil {
		panic(err)
	}

	log.Printf("Successfully migrated database")
}

func Down(cfg config.Config, migrationsDir fs.FS) {
	db, _ := sql.Open("pgx", cfg.PostgresURL)
	if err := migrate.Down(context.Background(), db, migrationsDir); err != nil {
		panic(err)
	}
}
