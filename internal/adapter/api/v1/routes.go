package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcelofabianov/picpay/pkg/zap"
)

func SetupRoutes(router *fiber.Router, logger *zap.Logger) {
	group := (*router).Group("/v1")

	// Users
	group.Get("/users", GetUsersHandler)
	group.Get("/users/:id", GetUserHandler)
	group.Post("/users", CreateUserHandler)
	group.Put("/users/:id", UpdateUserHandler)
	group.Delete("/users/:id", DeleteUserHandler)
}
