package main

import (
	"log"

	"github.com/MetaDandy/maquetador-angular-backend/config"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config.Load()

	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, API")
	})

	log.Println("Server started on port" + config.Port)
	app.Listen(":" + config.Port)

}
