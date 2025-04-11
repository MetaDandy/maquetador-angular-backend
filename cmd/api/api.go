package api

import (
	"github.com/MetaDandy/maquetador-angular-backend/src"
	"github.com/gofiber/fiber/v2"
)

func SetupApi(app *fiber.App, c *src.Container) {
	api := app.Group("/api")

	handlers := []func(fiber.Router){
		c.UserHandler.RegisterUserRoutes,
		c.ProjectHandler.RegisterProjectRoutes,
	}

	for _, register := range handlers {
		register(api)
	}
}
