package main

import (
	"log"
	"os"

	"github.com/MetaDandy/maquetador-angular-backend/cmd/api"
	"github.com/MetaDandy/maquetador-angular-backend/config"
	"github.com/MetaDandy/maquetador-angular-backend/middleware"
	"github.com/MetaDandy/maquetador-angular-backend/src"
	"github.com/MetaDandy/maquetador-angular-backend/src/modules/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.Load()

	app := fiber.New()
	app.Use(middleware.Logger())

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("ALLOW_ORIGINS"),
		AllowMethods: "GET,POST,PATCH,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Aloha")
	})

	websocket.InitializeWebsocket(app)

	c := src.SetupContainer()
	api.SetupApi(app, c)

	log.Println("Server started on port" + config.Port)
	app.Listen(":" + config.Port)

}
