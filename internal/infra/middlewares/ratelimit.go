package middlewares

import "github.com/gofiber/fiber/v2"

func RateLimitMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//...

		return c.Next()
	}
}
