package app

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"news-api/internal/models"
)

// HandleGetAllNews обработчик для получения всех новостей постранично
func (a *App) HandleGetAllNews(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	newsList, err := a.DB.GetAllNewsWithPagination(limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch news"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "news": newsList})
}

// HandleGetNewsByID обработчик для получения новости по id
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

// HandleCreateNews обработчик для создания новостей
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

// HandleUpdateNews обработчик для обновления новостей
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

// HandleEditNews обработчик для обновления новостей с категориями
func (a *App) HandleEditNews(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID"})
	}

	var input struct {
		Title      *string `json:"Title"`
		Content    *string `json:"Content"`
		Categories []int64 `json:"Categories"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	news, err := a.DB.GetNewsByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "news not found"})
	}

	if input.Title != nil {
		news.Title = *input.Title
	}
	if input.Content != nil {
		news.Content = *input.Content
	}

	if err := a.DB.UpdateNews(news); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update news"})
	}

	if input.Categories != nil {
		if err := a.DB.UpdateNewsCategories(id, input.Categories); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update categories"})
		}
	}

	return c.JSON(http.StatusOK, news)
}

// HandleDeleteNews обработчик для удаления новостей
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

// GetAllCategories обработчик для получения всех категорий
func (a *App) GetAllCategories(c echo.Context) error {
	categories, err := a.DB.GetAllCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch categories"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "categories": categories})
}

// GetCategoryByID обработчик для получения категории по id
func (a *App) GetCategoryByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID"})
	}
	category, err := a.DB.GetCategoryByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "category not found"})
	}
	return c.JSON(http.StatusOK, category)
}

// CreateCategory обработчик для создания категории
func (a *App) CreateCategory(c echo.Context) error {
	var category models.Category
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}
	if err := a.DB.CreateCategory(&category); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create category"})
	}
	return c.JSON(http.StatusCreated, category)
}

// UpdateCategory обработчик для обновления категории
func (a *App) UpdateCategory(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID"})
	}
	var category models.Category
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}
	category.ID = id
	if err := a.DB.UpdateCategory(&category); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update category"})
	}
	return c.JSON(http.StatusOK, category)
}

// DeleteCategory обработчик для удаления категории
func (a *App) DeleteCategory(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID"})
	}
	if err := a.DB.DeleteCategory(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete category"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "category deleted"})
}
