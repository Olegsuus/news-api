package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"news-api/internal/app"
	"news-api/internal/config"
	"news-api/internal/database"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewDB(cfg)
	migrations.Migrations(cfg, db.DB)

	app := &app.App{
		Config: cfg,
		DB:     db,
		Echo:   fiber.New(),
	}

	srv := &app.Server{}
	app.ServerInterface = srv

	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
