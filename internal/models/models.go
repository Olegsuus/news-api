package models

//go:generate reform
//reform:news
type News struct {
	ID         int64  `reform:"id,pk" json:"id"`
	Title      string `reform:"title" json:"title"`
	Content    string `reform:"content" json:"content"`
	Categories []int  `reform:"-" json:"categories"`
}

//go:generate reform
//reforms:categories
type NewsCategory struct {
	NewsID     int64 `reform:"news_id,pk" json:"news_id"`
	CategoryID int64 `reform:"category_id,pk" json:"category_id"`
}
