package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"news-api/internal/config"
	"news-api/internal/models"
)

type App struct {
	Config          *config.Config
	DB              Storage
	ServerInterface ServerInterface
	Echo            *fiber.App
}

type Storage interface {
	Stop() error
	GetAllNews() ([]models.News, error)
	GetNewsByID(int64) (*models.News, error)
	CreateNews(*models.News) error
	UpdateNews(*models.News) error
	DeleteNews(int64) error
}

func (a *App) Start() error {
	a.ServerInterface.GetServer(a)

	addr := fmt.Sprintf(":%d", a.Config.Server.Port)
	log.Printf("Starting server on %s", addr)
	return a.Echo.Listen(addr)
}

// Stop закрывает соединение с базой данных, если есть ошибки
func (a *App) Stop() {
	if err := a.DB.Stop(); err != nil {
		log.Fatalf("Failed to close database: %v", err)
	}
}
