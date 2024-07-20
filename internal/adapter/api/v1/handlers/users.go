package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/marcelofabianov/picpay/internal/adapter/api/v1/presenter"
	"github.com/marcelofabianov/picpay/internal/adapter/api/v1/requests"
	"github.com/marcelofabianov/picpay/internal/infra/request"
	"github.com/marcelofabianov/picpay/internal/infra/response"
)

func GetUsersHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   []string{"user1", "user2"},
	})
}

func GetUserHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	id := c.Params("id")

	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   "user: " + id,
	})
}

func CreateUserHandler(c *fiber.Ctx) error {
	var data requests.UserCreateRequest
	c.BodyParser(&data)

	result := request.IsValid(c, data)
	if result {
		response.Created(c, presenter.UserPresenter{
			ID:               request.GenerateUUID(),
			Name:             data.Name,
			Email:            data.Email,
			DocumentRegistry: data.DocumentRegistry,
			CreatedAt:        request.GetCurrentTime(),
			UpdatedAt:        request.GetCurrentTime(),
			Enabled:          true,
		})
	}

	return nil
}

func UpdateUserHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	id := c.Params("id")

	return c.JSON(fiber.Map{
		"status": "OK",
		"data: ": "user updated: " + id,
	})
}

func DeleteUserHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	id := c.Params("id")

	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   "user deleted: " + id,
	})
}
