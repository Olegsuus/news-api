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
	Config          *config.Config
	DB              *database.DB
	ServerInterface ServerInterface
	Echo            *fiber.App
}

type ServerInterface interface {
	GetServer(*App)
}

func (a *App) Start() error {
	a.Echo.Use(logger.New())

	a.ServerInterface.GetServer(a)

	addr := fmt.Sprintf(":%d", a.Config.Server.Port)
	log.Printf("Starting server on %s", addr)
	return a.Echo.Listen(addr)
}

func (a *App) Stop() {
	if err := a.DB.Close(); err != nil {
		log.Fatalf("Failed to close database: %v", err)
	}
}
