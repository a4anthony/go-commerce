package router

import (
	"github.com/a4anthony/go-commerce/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", handlers.HandleHealthCheck)

	// Public Routes
	app.Post("/register", handlers.Register)
}
