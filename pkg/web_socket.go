package pkg

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

func WebSocket(c *websocket.Conn) {
	defer func() {
		log.Println("Desconectado", c.LocalAddr())
		c.Close()
	}()

	log.Println("Cliente conectado", c.LocalAddr())

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("Error al leer el mensaje", err)
			break
		}

		log.Printf("Mensaje recibido: %s", msg)

		if err := c.WriteMessage(mt, msg); err != nil {
			log.Println("Erro enviando el mensaje: ", err)
			break
		}
	}
}
