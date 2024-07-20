package middlewares

import "github.com/gofiber/fiber/v2"

func CorsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//...

		return c.Next()
	}
}
