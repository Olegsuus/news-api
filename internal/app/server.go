package app

type ServerInterface interface {
	GetServer(*App)
}

type Server struct {
}

// GetServer метод для запуска роутера и обработчика запросов
func (s *Server) GetServer(app *App) {

}
