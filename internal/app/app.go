package app

import (
	"fmt"
	_ "fmt"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/labstack/echo/v4/middleware"
	"log"
	"news-api/internal/models"

	"news-api/internal/config"

	"github.com/labstack/echo/v4"
	_ "news-api/internal/models"
)

type App struct {
	Config          *config.Config
	DB              Storage
	ServerInterface ServerInterface
	Echo            *echo.Echo
}

type Storage interface {
	Stop() error
	GetAllNews() ([]models.News, error)
	GetNewsByID(int64) (*models.News, error)
	CreateNews(*models.News) error
	DeleteNews(int64) error
	UpdateNews(*models.News) error
}

func (a *App) Start() error {
	a.Echo = echo.New()
	a.ServerInterface.GetServer(a)
	a.Echo.Use(middleware.Logger())
	a.Echo.Use(middleware.Recover())

	addr := fmt.Sprintf(":%d", a.Config.Server.Port)
	log.Printf("Starting server on %s", addr)
	return a.Echo.Start(addr)
}

// Stop закрывает если есть ошибки
func (a *App) Stop() {
	if err := a.DB.Stop(); err != nil {
		log.Fatalf("Failed to close database: %v", err)
	}
}
