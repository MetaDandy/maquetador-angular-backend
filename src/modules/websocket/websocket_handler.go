package websocket

import (
	"github.com/MetaDandy/maquetador-angular-backend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func InitializeWebsocket(app *fiber.App) {
	hub := NewHub()
	go hub.Run()

	app.Use("/api/v1/ws/:pageID", middleware.WebSocket())
	app.Get("/api/v1/ws/:pageID", websocket.New(WebSocket(hub)))
}
