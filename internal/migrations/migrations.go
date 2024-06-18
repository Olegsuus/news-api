package migrations

import (
	"embed"
	"log"

	"database/sql"
	"github.com/pressly/goose/v3"
	"news-api/internal/config"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func Migrations(cfg *config.Config, db *sql.DB) {

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set goose dialect: %v", err)
	}

	log.Println("Running migrations...")

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}
