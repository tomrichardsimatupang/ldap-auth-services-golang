package main

import (
	"ldap/internal/routes"
	"ldap/pkg/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config, err := utils.GetConfig()

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("config", config)
		return c.Next()
	})

	routes.Routes(app)

	log.Fatal(app.Listen(":3000"))

}
