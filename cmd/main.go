package main

import (
	"log"

	"github.com/MetaDandy/maquetador-angular-backend/cmd/api"
	"github.com/MetaDandy/maquetador-angular-backend/config"
	"github.com/MetaDandy/maquetador-angular-backend/middleware"
	"github.com/MetaDandy/maquetador-angular-backend/pkg"
	"github.com/MetaDandy/maquetador-angular-backend/src"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	config.Load()

	app := fiber.New()
	app.Use(middleware.Logger())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, API")
	})

	// Middleware
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(pkg.WebSocket))

	c := src.SetupContainer()
	api.SetupApi(app, c)

	log.Println("Server started on port" + config.Port)
	app.Listen(":" + config.Port)

}
