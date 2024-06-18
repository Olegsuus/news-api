package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"news-api/internal/database"
	"news-api/internal/models"
)

var validate = validator.New()

type EditNewsInput struct {
	Title      *string `json:"Title" validate:"omitempty,min=1"`
	Content    *string `json:"Content" validate:"omitempty,min=1"`
	Categories []int64 `json:"Categories"`
}

func EditNewsHandler(c echo.Context) error {
	id := c.Param("id")

	var input EditNewsInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	if err := validate.Struct(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "validation failed", "details": err.Error()})
	}

	news, err := database.DB.FindByPrimaryKeyFrom(models.NewsTable, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "news not found"})
	}

	if input.Title != nil {
		news.(*models.News).Title = *input.Title
	}
	if input.Content != nil {
		news.(*models.News).Content = *input.Content
	}

	if err := database.DB.Update(news); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update news"})
	}

	if input.Categories != nil {
		if _, err := database.DB.Exec("DELETE FROM NewsCategories WHERE news_id = $1", id); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update categories"})
		}
		for _, catID := range input.Categories {
			if _, err := database.DB.Exec("INSERT INTO NewsCategories (news_id, category_id) VALUES ($1, $2)", id, catID); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update categories"})
			}
		}
	}

	return c.JSON(http.StatusOK, news)
}

func ListNewsHandler(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	var newsList []models.News
	query := "SELECT id, title, content FROM News ORDER BY id LIMIT $1 OFFSET $2"
	rows, err := database.DB.Query(query, limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch news"})
	}
	defer rows.Close()

	for rows.Next() {
		var news models.News
		if err := rows.Scan(&news.ID, &news.Title, &news.Content); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to scan news"})
		}
		newsList = append(newsList, news)
	}

	var response struct {
		Success bool          `json:"Success"`
		News    []models.News `json:"News"`
	}
	response.Success = true
	response.News = newsList

	return c.JSON(http.StatusOK, response)
}
