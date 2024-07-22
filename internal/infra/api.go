package infra

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"

	"github.com/marcelofabianov/picpay/config"
	"github.com/marcelofabianov/picpay/internal"
	v1 "github.com/marcelofabianov/picpay/internal/adapter/api/v1"
	"github.com/marcelofabianov/picpay/internal/infra/middlewares"
	"github.com/marcelofabianov/picpay/internal/port"
	"github.com/marcelofabianov/picpay/pkg/zap"
)

func Api(cfg *config.Config, logger *zap.Logger, db *pgx.Conn) *fiber.App {
	app := fiber.New()

	app = ContainerInjection(app, db)

	api := app.Group("/api")

	if cfg.Api.Logging {
		api.Use(middlewares.LoggingMiddleware(logger))
	}

	api.Use(middlewares.RateLimitMiddleware())
	api.Use(middlewares.CorsMiddleware())

	api.Get("/health", func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		return c.JSON(fiber.Map{
			"status":  "OK",
			"message": "Service is healthy",
		})
	})

	v1.SetupRoutes(&api, logger)

	return app
}

func ContainerInjection(app *fiber.App, db *pgx.Conn) *fiber.App {
	app.Use(func(c *fiber.Ctx) error {
		container := internal.NewContainer(db)

		var userService port.UserService
		err := container.Invoke(func(service port.UserService) {
			userService = service
		})
		if err != nil {
			return err
		}

		c.Locals("userService", userService)

		return c.Next()
	})

	return app
}
