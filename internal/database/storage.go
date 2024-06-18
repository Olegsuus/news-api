package database

import (
	"fmt"
	"log"

	"database/sql"
	_ "github.com/lib/pq"
	"news-api/internal/config"
)

type DataBase struct {
	DB *sql.DB
}

// GetStorage функция для подключения к Базе Данных
func (db *DataBase) GetStorage(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

	var err error
	db.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}

// Stop метод для закрытие БД
func (db *DataBase) Stop() error {
	if db.DB != nil {
		err := db.DB.Close()
		{
			if err != nil {
				log.Fatalf("Failed to closed database: %v", err)
				return err
			}
		}
	}
	return nil
}
