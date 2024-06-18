package migrations

import (
	"database/sql"
	"embed"
	"github.com/pressly/goose/v3"
	"log"
)

//go:embed *.sql
var embedMigrations embed.FS

func MigrateUp(db *sql.DB) {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set goose dialect: %v", err)
	}

	if err := goose.Up(db, "."); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Success migrations")
}
