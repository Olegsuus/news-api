package app

type ServerInterface interface {
	GetServer(*App)
}

type Server struct{}

func (s *Server) GetServer(app *App) {
	app.Echo.Get("/list", app.HandleGetAllNews)
	app.Echo.Post("/edit/:id", app.HandleUpdateNews)
}
