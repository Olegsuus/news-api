package app

type ServerInterface interface {
	GetServer(*App)
}

type Server struct {
}

// GetServer метод для запуска роутера и обработчика запросов
func (s *Server) GetServer(app *App) {
	app.Echo.GET("/news", app.HandleGetAllNews)
	app.Echo.GET("/news/:id", app.HandleGetNewsByID)
	app.Echo.POST("/news", app.HandleCreateNews)
	app.Echo.PUT("/news/:id", app.HandleUpdateNews)
	app.Echo.DELETE("/news/:id", app.HandleDeleteNews)
}
