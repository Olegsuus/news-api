package main

import (
	"github.com/labstack/echo/v4"
	"news-api/internal/config"
	"news-api/internal/database"
	"news-api/internal/handlers"
)

func main() {
	// Инициализация конфигурации
	cfg := config.InitConfig()

	// Инициализация базы данных
	database.InitDB(cfg)

	// Создание нового Echo инстанса
	e := echo.New()

	// Регистрация маршрутов
	e.GET("/news", handlers.ListNewsHandler)
	e.PUT("/news/:id", handlers.EditNewsHandler)

	// Запуск сервера
	e.Start(":8090")
}
