package app

import (
	"fmt"
	"log"
	"news-api/internal/config"
	"news-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	Config          *config.Config
	DB              Storage
	ServerInterface ServerInterface
	Fiber           *fiber.App
}

type Storage interface {
	Stop() error
	GetAllNews() ([]models.News, error)
	GetNewsByID(int64) (*models.News, error)
	UpdateNews(*models.News) error
}

func (a *App) Start() error {
	a.ServerInterface.GetServer(a)
	addr := fmt.Sprintf(":%d", a.Config.Server.Port)
	log.Printf("Starting server on %s", addr)
	return a.Fiber.Listen(addr)
}

// Stop закрывает соединение с базой данных
func (a *App) Stop() {
	if err := a.DB.Stop(); err != nil {
		log.Fatalf("Failed to close database: %v", err)
	}
}
