package app

type ServerInterface interface {
	GetServer(*App)
}

type Server struct{}

func (s *Server) GetServer(app *App) {
	app.Fiber.Post("/edit/:id", app.HandleEditNews)
	app.Fiber.Get("/list", app.HandleListNews)
}
