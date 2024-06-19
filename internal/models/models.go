package models

//go:generate reform
//reform:news
type News struct {
	ID      int64  `reform:"id,pk" json:"id"`
	Title   string `reform:"title" json:"title"`
	Content string `reform:"content" json:"content"`
}

// NewsCategory представляет связь между новостью и категорией.
type NewsCategory struct {
	NewsID     int64 `reform:"news_id,pk" json:"news_id"`
	CategoryID int64 `reform:"category_id,pk" json:"category_id"`
}

//// Implementing reform.View interface
//func (*News) Schema() string {
//	return "public"
//}
//
//func (*News) String() string {
//	return "News"
//}
//
//// Implementing reform.Record interface
//func (n *News) Values() []interface{} {
//	return []interface{}{n.ID, n.Title, n.Content}
//}
//
//func (n *News) PKPointer() interface{} {
//	return &n.ID
//}

//// Implementing reform.View interface
//func (*NewsCategory) Schema() string {
//	return "public"
//}
//
//func (*NewsCategory) String() string {
//	return "NewsCategory"
//}
//
//// Implementing reform.Record interface
//func (nc *NewsCategory) Values() []interface{} {
//	return []interface{}{nc.NewsID, nc.CategoryID}
//}
//
//func (nc *NewsCategory) PKPointer() interface{} {
//	return &nc.NewsID
//}
