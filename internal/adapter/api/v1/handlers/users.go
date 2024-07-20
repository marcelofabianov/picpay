package handlers

import "github.com/gofiber/fiber/v2"

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
	c.Accepts("application/json")

	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   "user created",
	})
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
