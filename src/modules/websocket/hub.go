package websocket

import "log"

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast: make(chan struct {
			pageID string
			data   []byte
		}),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			if h.clients[client.pageID] == nil {
				h.clients[client.pageID] = make(map[*Client]bool)
			}
			h.clients[client.pageID][client] = true
			h.mu.Unlock()
		case client := <-h.unregister:
			h.mu.Unlock()
			if conns := h.clients[client.pageID]; conns != nil {
				delete(conns, client)
				close(client.sendChan)
				if len(conns) == 0 {
					delete(h.clients, client.pageID)
				}
			}
			h.mu.Unlock()
		case msg := <-h.broadcast:
			h.mu.Lock()
			for client := range h.clients[msg.pageID] {
				log.Printf("Enviando mensaje a página %s", msg.pageID)
				select {
				case client.sendChan <- msg.data:
				default:
					close(client.sendChan)
					delete(h.clients[msg.pageID], client)
				}
			}
			h.mu.Unlock()
		}
	}
}
