package migrations

import (
	"database/sql"
	"embed"
	"github.com/pressly/goose/v3"
	"log"
	"news-api/internal/config"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func Migrations(cfg *config.Config, db *sql.DB) {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect(cfg.Database.Driver); err != nil {
		log.Fatalf("Failed to set goose dialect: %v", err)
	}

	if err := goose.Up(db, cfg.Database.MigrationPath); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Success migrations")
}
