package models

//go:generate reform
//reform:news
type News struct {
	ID         int64   `reform:"id,pk" json:"id"`
	Title      string  `reform:"title" json:"title"`
	Content    string  `reform:"content" json:"content"`
	Categories []int64 `reform:"-" json:"categories"`
}

//go:generate reform
//reform:newscategories
type NewsCategory struct {
	NewsID     int64 `reform:"news_id,pk" json:"news_id"`
	CategoryID int64 `reform:"category_id " json:"category_id"`
}

//go:generate reform
//reform:categories
type Category struct {
	ID   int64  `reform:"id,pk" json:"id"`
	Name string `reform:"name" json:"name"`
}
