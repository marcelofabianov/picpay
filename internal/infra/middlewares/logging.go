package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/marcelofabianov/picpay/pkg/zap"
)

func LoggingMiddleware(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		logger.Info(
			"Request completed",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("remote_addr", c.IP()),
			zap.String("duration", time.Since(start).String()),
			zap.String("version", "v1"),
			zap.Int("status", c.Response().StatusCode()),
			zap.Error(err),
		)
		return err
	}
}
