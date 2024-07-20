package infra

import (
	"github.com/gofiber/fiber/v2"

	"github.com/marcelofabianov/picpay/config"
	"github.com/marcelofabianov/picpay/internal/adapter/api/common/middlewares"
	v1 "github.com/marcelofabianov/picpay/internal/adapter/api/v1/routes"
	"github.com/marcelofabianov/picpay/pkg/zap"
)

func Api(cfg *config.Config, logger *zap.Logger) *fiber.App {
	app := fiber.New()

	api := app.Group("/api")

	if cfg.Api.Logging {
		api.Use(middlewares.LoggingMiddleware(logger))
	}

	api.Use(middlewares.RateLimitMiddleware())

	api.Get("/health", HealthCheckHandler)

	v1.SetupRoutes(&api, logger)

	return app
}

func HealthCheckHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Service is healthy",
	})
}
