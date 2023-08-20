package handlers

import "github.com/gofiber/fiber/v2"

func GetAllCategories(c *fiber.Ctx) error {
	return c.SendString("All categories")
}
