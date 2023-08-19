package api

import (
	"github.com/gofiber/fiber/v2"
)

func V1Welcome(c *fiber.Ctx) error {

	json := fiber.Map{
		"message": "API Run Well",
		"status":  "success",
		"hello":   "yes",
	}

	return c.Status(fiber.StatusCreated).JSON(json)

}
