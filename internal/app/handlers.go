package app

import (
	"news-api/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// HandleEditNews обработчик для обновления новости
func (a *App) HandleEditNews(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	var input models.News
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	input.ID = id
	if err := a.DB.UpdateNews(&input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update news"})
	}

	return c.JSON(fiber.Map{"success": true, "news": input})
}

// HandleListNews обработчик для получения списка новостей
func (a *App) HandleListNews(c *fiber.Ctx) error {
	newsList, err := a.DB.GetAllNews()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch news"})
	}

	return c.JSON(fiber.Map{"success": true, "news": newsList})
}
