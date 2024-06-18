package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"news-api/internal/handlers"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/edit/:id", handlers.EditNewsHandler)
	e.GET("/list", handlers.ListNewsHandler)

	return e
}
