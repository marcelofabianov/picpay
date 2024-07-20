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

func BadRequestErrorsResponse(c *fiber.Ctx, errors []map[string]string) {
	c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
		Success: false,
		Status:  400,
		Error:   errors,
		Data:    nil,
	})
}

func BadRequestResponse(c *fiber.Ctx, err error) {
	c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
		Success: false,
		Status:  400,
		Error:   []map[string]string{{"error": err.Error()}},
		Data:    nil,
	})
}
