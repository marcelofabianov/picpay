package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcelofabianov/picpay/internal/infra/request"
	"github.com/marcelofabianov/picpay/internal/infra/response"
	"github.com/marcelofabianov/picpay/internal/port"
)

func GetUsersHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var users []port.UserPresenter

	users = append(users, port.UserPresenter{
		ID:               request.GenerateUUID(),
		Name:             "John Doe",
		Email:            "john@email.com",
		DocumentRegistry: "11111111111",
		CreatedAt:        request.GetCurrentTime(),
		UpdatedAt:        request.GetCurrentTime(),
		Enabled:          true,
	})

	users = append(users, port.UserPresenter{
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

	response.Ok(c, port.UserPresenter{
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
	var data port.UserCreateRequest
	c.BodyParser(&data)

	result := request.IsValid(c, data)
	if result {
		inputService := port.UserCreateRequest{
			Name:             data.Name,
			Email:            data.Email,
			Password:         data.Password,
			DocumentRegistry: data.DocumentRegistry,
		}

		service := c.Locals("userService").(port.UserService)

		outputService, err := service.CreateUser(c.Context(), inputService)
		if err != nil {
			response.InternalServerError(c)
		}

		presenter := port.UserPresenter{
			ID:               outputService.ID,
			Name:             outputService.Name,
			Email:            outputService.Email,
			DocumentRegistry: outputService.DocumentRegistry,
			CreatedAt:        outputService.CreatedAt,
			UpdatedAt:        outputService.UpdatedAt,
			Enabled:          outputService.Enabled,
		}

		response.Created(c, presenter)
	}

	return nil
}

func UpdateUserHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	id := c.Params("id")

	response.Ok(c, port.UserPresenter{
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
