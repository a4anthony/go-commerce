package router

import (
	"fmt"

	"github.com/a4anthony/go-commerce/handlers"
	"github.com/a4anthony/go-commerce/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", handlers.HandleHealthCheck)

	api := app.Group("/api")

	// private := api.Use(middlewares.JwtAuthMiddleware)
	usersGroup := api.Group("/users")

	usersGroupPrivate := api.Group("/users")

	fmt.Println(usersGroupPrivate)

	usersGroup.Post("/register", handlers.Register)
	usersGroup.Post("/login", handlers.Login)
	usersGroupPrivate.Get("/me", middlewares.JwtAuthMiddleware(), handlers.Me)
}
