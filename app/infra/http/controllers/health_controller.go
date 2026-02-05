package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleHealth(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Server is running",
	})
}
