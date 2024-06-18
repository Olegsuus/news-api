package models

import (
	"fmt"
	"time"
)

// News представляет структуру новости.
type News struct {
	ID        int64     `reform:"id,pk" json:"id"`
	Title     string    `reform:"title" json:"title"`
	Content   string    `reform:"content" json:"content"`
	CreatedAt time.Time `reform:"created_at" json:"created_at"`
	UpdatedAt time.Time `reform:"updated_at" json:"updated_at"`
}

// String implements reform.Record interface.
func (n *News) String() string {
	return fmt.Sprintf("News ID: %d, Title: %s", n.ID, n.Title)
}

// StringValues implements reform.Record interface.
func (n *News) StringValues() []interface{} {
	return []interface{}{n.ID, n.Title, n.Content, n.CreatedAt, n.UpdatedAt}
}

// Schema implements reform.View interface.
func (n *News) Schema() string {
	return NewsTable
}

// StringName implements reform.View interface.
func (n *News) StringName() string {
	return "News"
}

// NewsTable - имя таблицы
const NewsTable = "news"
