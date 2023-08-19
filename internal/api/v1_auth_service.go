package api

import (
	"ldap/pkg/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func V1AuthService(c *fiber.Ctx) error {

	config := c.Locals("config").(utils.Config)

	user, err := utils.GetUserLDAP(config.Ldap, "tom.richard", "P@ssw0rd")

	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	var authRequest AuthRequest

	if err := c.BodyParser(&authRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Bad request"})
	}

	json := fiber.Map{
		"uid":   user.Uid,
		"name":  user.Name,
		"mail":  user.Mail,
		"phone": user.Phone,
	}

	return c.Status(fiber.StatusCreated).JSON(json)

}
