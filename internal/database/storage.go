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

// GetAllNewsWithPagination метод для получения всех новостей через пагинацию
func (db *DataBase) GetAllNewsWithPagination(limit, offset int) ([]models.News, error) {
	rows, err := db.DB.Query("SELECT id, title, content FROM News ORDER BY id LIMIT $1 OFFSET $2", limit, offset)
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

// GetNewsByID метод для получения новости по id
func (db *DataBase) GetNewsByID(id int64) (*models.News, error) {
	var news models.News
	err := db.DB.QueryRow("SELECT id, title, content FROM News WHERE id = $1", id).Scan(&news.ID, &news.Title, &news.Content)
	if err != nil {
		return nil, err
	}
	return &news, nil
}

// CreateNews метод для создания новости
func (db *DataBase) CreateNews(news *models.News) error {
	_, err := db.DB.Exec("INSERT INTO News (title, content) VALUES ($1, $2)", news.Title, news.Content)
	return err
}

// UpdateNews метод для обновления новости
func (db *DataBase) UpdateNews(news *models.News) error {
	_, err := db.DB.Exec("UPDATE News SET title = $1, content = $2 WHERE id = $3", news.Title, news.Content, news.ID)
	return err
}

// DeleteNews метод для удаления новости
func (db *DataBase) DeleteNews(id int64) error {
	_, err := db.DB.Exec("DELETE FROM News WHERE id = $1", id)
	return err
}

func (db *DataBase) UpdateNewsCategories(newsID int64, categories []int64) error {
	// Удалить старые категории
	_, err := db.DB.Exec("DELETE FROM NewsCategories WHERE news_id = $1", newsID)
	if err != nil {
		return err
	}
	// Добавить новые категории
	for _, categoryID := range categories {
		_, err := db.DB.Exec("INSERT INTO NewsCategories (news_id, category_id) VALUES ($1, $2)", newsID, categoryID)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAllCategories метод для получения всех категорий
func (db *DataBase) GetAllCategories() ([]models.Category, error) {
	rows, err := db.DB.Query("SELECT id, name FROM Categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

// GetCategoryByID метод для получения категорий по id
func (db *DataBase) GetCategoryByID(id int64) (*models.Category, error) {
	var category models.Category
	err := db.DB.QueryRow("SELECT id, name FROM Categories WHERE id = $1", id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// CreateCategory метод для добавления категорий
func (db *DataBase) CreateCategory(category *models.Category) error {
	_, err := db.DB.Exec("INSERT INTO Categories (name) VALUES ($1)", category.Name)
	return err
}

func (db *DataBase) UpdateCategory(category *models.Category) error {
	_, err := db.DB.Exec("UPDATE Categories SET name = $1 WHERE id = $2", category.Name, category.ID)
	return err
}

// DeleteCategory метод для удаления категории
func (db *DataBase) DeleteCategory(id int64) error {
	_, err := db.DB.Exec("DELETE FROM Categories WHERE id = $1", id)
	return err
}
