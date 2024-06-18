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

type DB struct {
	DB *reform.DB
}

func NewDB(cfg *config.Config) *DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	db := reform.NewDB(sqlDB, postgresql.Dialect, nil)
	return &DB{DB: db}
}

func (db *DB) Close() error {
	return db.DB.DBInterface().(*sql.DB).Close()
}

func (db *DB) GetAllNews(limit, offset int) ([]models.News, error) {
	var newsList []models.News
	err := db.DB.SelectAllFrom(models.NewsTable, "LIMIT ? OFFSET ?", limit, offset, &newsList)
	return newsList, err
}

func (db *DB) GetNewsByID(id int64) (*models.News, error) {
	news := &models.News{}
	err := db.DB.FindByPrimaryKeyTo(news, id)
	return news, err
}

func (db *DB) CreateNews(news *models.News) error {
	return db.DB.Save(news)
}

func (db *DB) UpdateNews(news *models.News) error {
	return db.DB.Update(news)
}

func (db *DB) DeleteNews(id int64) error {
	news := &models.News{ID: id}
	return db.DB.Delete(news)
}
