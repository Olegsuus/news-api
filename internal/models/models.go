package models

// News представляет структуру новости.
type News struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// NewsCategory представляет связь между новостью и категорией.
type NewsCategory struct {
	NewsID     int64 `json:"news_id"`
	CategoryID int64 `json:"category_id"`
}

// Category представляет структуру категории.
type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
