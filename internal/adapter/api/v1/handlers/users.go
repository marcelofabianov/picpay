package handlers

import "github.com/gofiber/fiber/v2"

func GetUsersHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   []string{"user1", "user2"},
	})
}
