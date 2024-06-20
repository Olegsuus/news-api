package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"news-api/internal/config"
	"news-api/internal/database"
)

type App struct {
	Config *config.Config
	DB     *database.DB
	Fiber  *fiber.App
}

func (a *App) Start() error {
	a.Fiber.Use(logger.New())
	a.Fiber.Use(BasicAuthorization)

	a.Fiber.Get("/list", a.HandleGetAllNews)

	a.Fiber.Post("/create", a.CreateNew)
	a.Fiber.Post("/edit/:id", a.HandleUpdateNews)

	addr := fmt.Sprintf(":%d", a.Config.Server.Port)
	log.Printf("Starting server on %s", addr)

	return a.Fiber.Listen(addr)
}

func (a *App) Stop() {
	if err := a.DB.Close(); err != nil {
		log.Fatalf("Failed to close database: %v", err)
	}
}
