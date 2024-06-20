package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"log"
	"news-api/internal/config"
	"news-api/internal/models"
)

type DB struct {
	DB *reform.DB
}

func NewDB(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	return sqlDB
}

func CreateReformDB(sqlDB *sql.DB) *DB {
	db := reform.NewDB(sqlDB, postgresql.Dialect, nil)
	return &DB{DB: db}
}

func (db *DB) Close() error {
	return db.DB.DBInterface().(*sql.DB).Close()
}

func (db *DB) GetAllNews(limit, offset int) ([]models.News, error) {
	var structs []reform.Struct

	query := fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)
	structs, err := db.DB.SelectAllFrom(models.NewsTable, query)
	if err != nil {
		return nil, err
	}

	newsList := make([]models.News, len(structs))
	for i, s := range structs {
		if news, ok := s.(*models.News); ok {
			categoryIDs := []int64{}
			query := fmt.Sprintf("WHERE news_id = %d", news.ID)
			newsCategoryStructs, err := db.DB.SelectAllFrom(models.NewsCategoryTable, query)
			if err != nil {
				log.Printf("Error selecting categories for news ID %d: %v", news.ID, err)
				continue
			}

			for _, ns := range newsCategoryStructs {
				if nc, ok := ns.(*models.NewsCategory); ok {
					categoryIDs = append(categoryIDs, nc.CategoryID)
				} else {
					log.Printf("Error parsing category ID for news ID %d: unexpected type %T", news.ID, ns)
				}
			}
			news.Categories = categoryIDs
			newsList[i] = *news
		} else {
			return nil, fmt.Errorf("unexpected type %T", s)
		}
	}
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

func (db *DB) CreateCategory(cat *models.Category) error {
	return db.DB.Save(cat)
}

func (db *DB) CreateNewsCategory(nc *models.NewsCategory) error {
	return db.DB.Save(nc)
}
