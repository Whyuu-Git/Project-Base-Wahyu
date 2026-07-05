package response

import "github.com/gofiber/fiber/v2"


func Success(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(data)
}


func Error(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "error",
		"message": message,
	})
}
