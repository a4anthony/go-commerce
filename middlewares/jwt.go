package middlewares

import (
	"log"

	"github.com/a4anthony/go-commerce/utils"
	"github.com/gofiber/fiber/v2"
)

// func JwtAuthMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("log middleware")
// 		err := utils.TokenValid(r)
// 		fmt.Println(err)

// 		if err != nil {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

func JwtAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("log middleware")
		err := utils.TokenValid(c)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		return c.Next()
	}
}
