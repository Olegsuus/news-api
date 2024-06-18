package database

import (
	"database/sql"
	"fmt"
	"log"
	"news-api/internal/config"
	"news-api/internal/models"

	_ "github.com/lib/pq"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

type DataBase struct {
	DB *reform.DB
}

func (db *DataBase) GetStorage(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	db.DB = reform.NewDB(sqlDB, postgresql.Dialect, nil)
}

func (db *DataBase) Stop() error {
	return db.DB.DBInterface().Close()
}

func (db *DataBase) GetAllNews() ([]models.News, error) {
	var newsList []models.News
	err := db.DB.FindAllFrom(models.NewsTable, &newsList)
	return newsList, err
}

func (db *DataBase) GetNewsByID(id int64) (*models.News, error) {
	news := &models.News{}
	err := db.DB.FindByPrimaryKeyTo(news, id)
	return news, err
}

func (db *DataBase) UpdateNews(news *models.News) error {
	return db.DB.Update(news)
}
