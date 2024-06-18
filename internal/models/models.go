package models

import (
	"time"
)

type News struct {
	ID        int64     `reform:"id,pk" json:"id"`
	Title     string    `reform:"title" json:"title"`
	Content   string    `reform:"content" json:"content"`
	CreatedAt time.Time `reform:"created_at" json:"created_at"`
	UpdatedAt time.Time `reform:"updated_at" json:"updated_at"`
}

// NewsTable is the reform table identifier for News.
const NewsTable = "news"

// NewsCategory represents the relationship between news and categories.
type NewsCategory struct {
	NewsID     int64 `reform:"news_id,pk" json:"news_id"`
	CategoryID int64 `reform:"category_id,pk" json:"category_id"`
}
