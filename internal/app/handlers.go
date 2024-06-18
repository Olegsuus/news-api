package app

import (
	"github.com/gofiber/fiber/v2"
	"news-api/internal/models"
	"strconv"
)

// HandleGetAllNews обработчик для получения всех новостей
func (a *App) HandleGetAllNews(c *fiber.Ctx) error {
	newsList, err := a.DB.GetAllNews()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch news"})
	}
	return c.JSON(fiber.Map{"success": true, "news": newsList})
}

// HandleGetNewsByID обработчик для получения новости по id
func (a *App) HandleGetNewsByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}
	news, err := a.DB.GetNewsByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "news not found"})
	}
	return c.JSON(news)
}

// HandleCreateNews обработчик для создания новостей
func (a *App) HandleCreateNews(c *fiber.Ctx) error {
	var news models.News
	if err := c.BodyParser(&news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}
	if err := a.DB.CreateNews(&news); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create news"})
	}
	return c.Status(fiber.StatusCreated).JSON(news)
}

// HandleUpdateNews обработчик для обновления новостей
func (a *App) HandleUpdateNews(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}
	var news models.News
	if err := c.BodyParser(&news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}
	news.ID = id
	if err := a.DB.UpdateNews(&news); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update news"})
	}
	return c.JSON(news)
}

// HandleDeleteNews обработчик для удаления новостей
func (a *App) HandleDeleteNews(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}
	if err := a.DB.DeleteNews(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete news"})
	}
	return c.JSON(fiber.Map{"message": "news deleted"})
}
