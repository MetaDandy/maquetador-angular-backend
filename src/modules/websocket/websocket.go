package websocket

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

func WebSocket(hub *Hub) func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		pageID := c.Params("pageID")
		client := &Client{
			conn:     c,
			pageID:   pageID,
			sendChan: make(chan []byte, 256),
		}

		hub.register <- client
		defer func() {
			hub.unregister <- client
			c.Close()
			log.Println("Desconectado", c.LocalAddr())
		}()

		log.Println("Cliente conectado", c.LocalAddr(), "en página", pageID)

		go func() {
			for msg := range client.sendChan {
				if err := c.WriteMessage(websocket.TextMessage, msg); err != nil {
					break
				}
			}
		}()

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("Error al leer:", err)
				break
			}
			log.Printf("Recibido de %s: %s", client.conn.LocalAddr(), msg)
			hub.broadcast <- struct {
				pageID string
				data   []byte
			}{pageID, msg}
		}
	}
}
