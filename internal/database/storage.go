package database

import (
	"fmt"
	"log"
	"news-api/internal/models"

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

func (db *DataBase) GetAllNews() ([]models.News, error) {
	rows, err := db.DB.Query("SELECT id, title, content FROM News")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var newsList []models.News
	for rows.Next() {
		var news models.News
		if err := rows.Scan(&news.ID, &news.Title, &news.Content); err != nil {
			return nil, err
		}
		newsList = append(newsList, news)
	}
	return newsList, nil
}

func (db *DataBase) GetNewsByID(id int64) (*models.News, error) {
	var news models.News
	err := db.DB.QueryRow("SELECT id, title, content FROM News WHERE id = $1", id).Scan(&news.ID, &news.Title, &news.Content)
	if err != nil {
		return nil, err
	}
	return &news, nil
}

func (db *DataBase) CreateNews(news *models.News) error {
	_, err := db.DB.Exec("INSERT INTO News (title, content) VALUES ($1, $2)", news.Title, news.Content)
	return err
}

func (db *DataBase) UpdateNews(news *models.News) error {
	_, err := db.DB.Exec("UPDATE News SET title = $1, content = $2 WHERE id = $3", news.Title, news.Content, news.ID)
	return err
}

func (db *DataBase) DeleteNews(id int64) error {
	_, err := db.DB.Exec("DELETE FROM News WHERE id = $1", id)
	return err
}
