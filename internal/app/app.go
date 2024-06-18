package app

import (
	"fmt"
	"log"

	"news-api/internal/config"
	"news-api/internal/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	Config          *config.Config
	DB              Storage
	ServerInterface ServerInterface
	Echo            *echo.Echo
}

type Storage interface {
	Stop() error
	GetAllNewsWithPagination(int, int) ([]models.News, error)
	GetNewsByID(int64) (*models.News, error)
	CreateNews(*models.News) error
	DeleteNews(int64) error
	UpdateNews(*models.News) error
	UpdateNewsCategories(int64, []int64) error
	DeleteCategory(int64) error
	UpdateCategory(*models.Category) error
	CreateCategory(*models.Category) error
	GetCategoryByID(int64) (*models.Category, error)
	GetAllCategories() ([]models.Category, error)
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
