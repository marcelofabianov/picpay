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

	var users []presenter.UserPresenter

	users = append(users, presenter.UserPresenter{
		ID:               request.GenerateUUID(),
		Name:             "John Doe",
		Email:            "john@email.com",
		DocumentRegistry: "11111111111",
		CreatedAt:        request.GetCurrentTime(),
		UpdatedAt:        request.GetCurrentTime(),
		Enabled:          true,
	})

	users = append(users, presenter.UserPresenter{
		ID:               request.GenerateUUID(),
		Name:             "Jane Doe",
		Email:            "jane@email.com",
		DocumentRegistry: "22222222222",
		CreatedAt:        request.GetCurrentTime(),
		UpdatedAt:        request.GetCurrentTime(),
		Enabled:          true,
	})

	response.Ok(c, users)

	return nil
}

func GetUserHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	id := c.Params("id")

	response.Ok(c, presenter.UserPresenter{
		ID:               id,
		Name:             "John Doe",
		Email:            "john@email.com",
		DocumentRegistry: "11111111111",
		CreatedAt:        request.GetCurrentTime(),
		UpdatedAt:        request.GetCurrentTime(),
		Enabled:          true,
	})

	return nil
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

	response.Ok(c, presenter.UserPresenter{
		ID:               id,
		Name:             "John",
		Email:            "john@email.com",
		DocumentRegistry: "11111111111",
		CreatedAt:        request.GetCurrentTime(),
		UpdatedAt:        request.GetCurrentTime(),
		Enabled:          true,
	})

	return nil
}

func DeleteUserHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	response.NoContent(c)

	return nil
}
