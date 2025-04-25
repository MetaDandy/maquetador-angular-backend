package websocket

import (
	"sync"

	"github.com/gofiber/websocket/v2"
)

type Client struct {
	conn     *websocket.Conn
	pageID   string
	sendChan chan []byte
}

type Hub struct {
	mu         sync.Mutex
	clients    map[string]map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan struct {
		pageID string
		data   []byte
	}
}
