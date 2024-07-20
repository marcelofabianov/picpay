package response

import "github.com/gofiber/fiber/v2"

type SuccessOkResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

func NewOkResponse(c *fiber.Ctx, data interface{}) {
	c.Status(fiber.StatusOK).JSON(SuccessOkResponse{
		Success: true,
		Status:  200,
		Error:   nil,
		Data:    data,
	})
}

func NewCreatedResponse(c *fiber.Ctx, data interface{}) {
	c.Status(fiber.StatusCreated).JSON(SuccessOkResponse{
		Success: true,
		Status:  201,
		Error:   nil,
		Data:    data,
	})
}

func NewNoContentResponse(c *fiber.Ctx) {
	c.Status(fiber.StatusNoContent)
}
