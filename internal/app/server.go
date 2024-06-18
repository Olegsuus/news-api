package app

type ServerInterface interface {
	GetServer(*App)
}

type Server struct {
}

// GetServer метод для запуска роутера и обработчика запросов
func (s *Server) GetServer(app *App) {
	app.Echo.GET("/news", app.HandleGetAllNews) // вывод производится по страницам
	app.Echo.GET("/news/:id", app.HandleGetNewsByID)
	app.Echo.POST("/news", app.HandleCreateNews)
	app.Echo.PUT("/news/:id", app.HandleUpdateNews)
	app.Echo.DELETE("/news/:id", app.HandleDeleteNews)
	app.Echo.POST("/edit/:id", app.HandleEditNews)

	app.Echo.GET("/categories", app.GetAllCategories)
	app.Echo.GET("/categories/:id", app.GetCategoryByID)
	app.Echo.POST("/categories", app.CreateCategory)
	app.Echo.PUT("/categories/:id", app.UpdateCategory)
	app.Echo.DELETE("/categories/:id", app.DeleteCategory)
}
