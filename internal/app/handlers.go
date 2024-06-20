package app

import (
	"github.com/gofiber/fiber/v2"
	"news-api/internal/models"
	"strconv"
)

func (a *App) HandleGetAllNews(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	newsList, err := a.DB.GetAllNews(limit, offset)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch news"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "news": newsList})
}

func (a *App) HandleGetNewsByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}
	news, err := a.DB.GetNewsByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "news not found"})
	}
	return c.Status(fiber.StatusOK).JSON(news)
}

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
	return c.Status(fiber.StatusOK).JSON(news)
}

func (a *App) HandleDeleteNews(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}
	if err := a.DB.DeleteNews(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete news"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "news deleted"})
}

func (a *App) CreateNew(c *fiber.Ctx) error {
	news := models.News{Title: "dfsfs", Content: "fdfs"}
	return a.DB.CreateNews(&news)
}
