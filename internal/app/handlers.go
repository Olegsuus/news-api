package app

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"news-api/internal/models"
)

func (a *App) HandleGetAllNews(c echo.Context) error {
	newsList, err := a.DB.GetAllNews()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch news"})
	}
	return c.JSON(http.StatusOK, newsList)
}

func (a *App) HandleGetNewsByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID"})
	}
	news, err := a.DB.GetNewsByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "news not found"})
	}
	return c.JSON(http.StatusOK, news)
}

func (a *App) HandleCreateNews(c echo.Context) error {
	var news models.News
	if err := c.Bind(&news); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}
	if err := a.DB.CreateNews(&news); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create news"})
	}
	return c.JSON(http.StatusCreated, news)
}

func (a *App) HandleUpdateNews(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID"})
	}
	var news models.News
	if err := c.Bind(&news); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}
	news.ID = id
	if err := a.DB.UpdateNews(&news); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update news"})
	}
	return c.JSON(http.StatusOK, news)
}

func (a *App) HandleDeleteNews(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID"})
	}
	if err := a.DB.DeleteNews(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete news"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "news deleted"})
}
