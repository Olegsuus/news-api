package app

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
)

const (
	username = "Oleg"
	password = "1234"
)

func BasicAuthorization(c *fiber.Ctx) error {
	// Получаем заголовок Authorization
	auth := c.Get("Authorization")
	if auth == "" {
		c.Status(fiber.StatusUnauthorized)
		return c.SendString("Unauthorized")
	}

	// Проверяем, что заголовок начинается с "Basic "
	if len(auth) < 6 || auth[:6] != "Basic " {
		c.Status(fiber.StatusUnauthorized)
		return c.SendString("Unauthorized")
	}

	// Декодируем базовые учетные данные из base64
	payload, err := base64.StdEncoding.DecodeString(auth[6:])
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.SendString("Unauthorized")
	}

	// Проверяем учетные данные
	if string(payload) != username+":"+password {
		c.Status(fiber.StatusUnauthorized)
		return c.SendString("Unauthorized")
	}

	// Продолжаем выполнение запроса
	return c.Next()
}
