package app

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type ServerInterface interface {
	GetServer(*App)
}

type Server struct {
}

// GetServer метод для запуска роутера и обработчика запросов
func (s *Server) GetServer(app *App) {
	app.Echo.Use(authMiddleware)

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

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing authorization header"})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != "1703" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
		}

		return next(c)
	}
}
