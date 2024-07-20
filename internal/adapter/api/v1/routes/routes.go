package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcelofabianov/picpay/internal/adapter/api/v1/handlers"
	"github.com/marcelofabianov/picpay/pkg/zap"
)

func SetupRoutes(router *fiber.Router, logger *zap.Logger) {
	group := (*router).Group("/v1")

	// Users
	group.Get("/users", handlers.GetUsersHandler)
	group.Get("/users/:id", handlers.GetUserHandler)
	group.Post("/users", handlers.CreateUserHandler)
	group.Put("/users/:id", handlers.UpdateUserHandler)
	group.Delete("/users/:id", handlers.DeleteUserHandler)
}
