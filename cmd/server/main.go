package main

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
	"news-api/internal/database"
	"news-api/internal/handlers"
)

func main() {
	// Инициализация конфигурации
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Инициализация базы данных
	database.InitDB()

	// Создание нового Echo инстанса
	e := echo.New()

	// Регистрация маршрутов
	e.GET("/news", handlers.ListNewsHandler)
	e.PUT("/news/:id", handlers.EditNewsHandler)

	// Запуск сервера
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
