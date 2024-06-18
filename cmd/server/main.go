package main

import (
	"log"
	migration "news-api/internal/migrations"

	"news-api/internal/app"
	"news-api/internal/config"
	"news-api/internal/database"
)

func main() {
	cfg := config.GetConfig()
	db := database.DataBase{}
	db.GetStorage(cfg)
	migration.Migrations(cfg, db.DB)
	App := app.App{Config: cfg, DB: &db}

	srv := &app.Server{}
	App.ServerInterface = srv

	if err := App.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
