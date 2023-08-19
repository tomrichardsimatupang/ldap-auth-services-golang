package routes

import (
	"ldap/internal/api"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	app.Get("/", api.V1Welcome)
	app.Post("/auth-service/v1/auth", api.V1AuthService)

}
