package response

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Success bool                `json:"success"`
	Status  int                 `json:"status"`
	Error   []map[string]string `json:"error"`
	Data    interface{}         `json:"data"`
}

func BadRequestErrors(c *fiber.Ctx, errors []map[string]string) {
	c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
		Success: false,
		Status:  400,
		Error:   errors,
		Data:    nil,
	})
}

func BadRequest(c *fiber.Ctx, err error) {
	c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
		Success: false,
		Status:  400,
		Error:   []map[string]string{{"error": err.Error()}},
		Data:    nil,
	})
}

func InternalServerError(c *fiber.Ctx) {
	c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
		Success: false,
		Status:  500,
		Error:   []map[string]string{{"error": "internal_server_error"}},
		Data:    nil,
	})
}
