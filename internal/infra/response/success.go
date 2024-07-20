package response

import "github.com/gofiber/fiber/v2"

type SuccessOkResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

func Ok(c *fiber.Ctx, data interface{}) {
	c.Status(fiber.StatusOK).JSON(SuccessOkResponse{
		Success: true,
		Status:  200,
		Error:   nil,
		Data:    data,
	})
}

func Created(c *fiber.Ctx, data interface{}) {
	c.Status(fiber.StatusCreated).JSON(SuccessOkResponse{
		Success: true,
		Status:  201,
		Error:   nil,
		Data:    data,
	})
}

func NoContent(c *fiber.Ctx) {
	c.Status(fiber.StatusNoContent)
}
