package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcelofabianov/picpay/config"
)

func NewApp(cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork: true,
		CaseSensitive: true,
		StrictRouting: false,
		ServerHeader: "Fiber",
		AppName: cfg.Name,
	})

	return app
}
