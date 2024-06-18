package models

// News представляет структуру новости.
type News struct {
	ID      int64  `db:"id" json:"id"`
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
}
