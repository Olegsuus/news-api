package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"news-api/internal/config"

	"github.com/spf13/viper"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

var DB *reform.DB

func InitDB(v *config.Config) {
	dsn := viper.GetString("database.dsn")
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	DB = reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(log.Printf))
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return DB.Query(query, args...)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return DB.Exec(query, args...)
}
