package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"news-api/internal/app"
	"news-api/internal/config"
	"news-api/internal/database"
	"news-api/internal/migrations"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewDB(cfg)
	migrations.Migrations(cfg, db)
	reformDB := database.CreateReformDB(db)

	app := &app.App{
		Config: cfg,
		DB:     reformDB,
		Fiber:  fiber.New(),
	}

	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
