package database

import (
	"fmt"
	"log"
	"news-api/internal/config"
	"news-api/internal/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DataBase struct {
	DB *sqlx.DB
}

func (db *DataBase) GetStorage(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

	sqlDB, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	db.DB = sqlDB
}

func (db *DataBase) Stop() error {
	return db.DB.Close()
}

func (db *DataBase) GetAllNews() ([]models.News, error) {
	var newsList []models.News
	err := db.DB.Select(&newsList, "SELECT * FROM news")
	return newsList, err
}

func (db *DataBase) GetNewsByID(id int64) (*models.News, error) {
	var news models.News
	err := db.DB.Get(&news, "SELECT * FROM news WHERE id=$1", id)
	return &news, err
}

func (db *DataBase) CreateNews(news *models.News) error {
	_, err := db.DB.NamedExec("INSERT INTO news (title, content) VALUES (:title, :content)", news)
	return err
}

func (db *DataBase) UpdateNews(news *models.News) error {
	_, err := db.DB.NamedExec("UPDATE news SET title=:title, content=:content WHERE id=:id", news)
	return err
}

func (db *DataBase) DeleteNews(id int64) error {
	_, err := db.DB.Exec("DELETE FROM news WHERE id=$1", id)
	return err
}
