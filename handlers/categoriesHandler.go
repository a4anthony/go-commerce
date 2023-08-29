package handlers

import (
	"github.com/a4anthony/go-commerce/database"
	"github.com/a4anthony/go-commerce/models"
	"github.com/gofiber/fiber/v2"
)

type CategoryResponse struct {
	Message  string          `json:"message"`
	Category models.Category `json:"category"`
}

func GetAllCategories(c *fiber.Ctx) error {
	categories := []models.Category{}

	database.DB.
		Preload("SubCategory").
		Preload("SubCategory.Category").
		Find(&categories)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Categories retrieved successfully",
		"categories": categories,
	})
}
