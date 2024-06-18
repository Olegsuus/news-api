package models

import _ "gopkg.in/reform.v1"

//go:generate reform

// reform:News
type News struct {
	ID      int64  `reform:"id,pk"`
	Title   string `reform:"title"`
	Content string `reform:"content"`
}

// reform:NewsCategories
type NewsCategories struct {
	NewsID     int64 `reform:"news_id"`
	CategoryID int64 `reform:"category_id"`
}
