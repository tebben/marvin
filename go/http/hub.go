package http

import (
	"encoding/json"
	"github.com/tebben/marvin/go/events"
	"github.com/tebben/marvin/go/marvin/models"
	"log"
)

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type Hub struct {
	// Registered connections.
	connections map[*Conn]bool

	// Inbound messages from the connections.
	broadcast chan []byte

	// Register requests from the connections.
	register chan *Conn

	// Unregister requests from connections.
	unregister chan *Conn
}

var hub = Hub{
	broadcast:   make(chan []byte),
	register:    make(chan *Conn),
	unregister:  make(chan *Conn),
	connections: make(map[*Conn]bool),
}

func (h *Hub) run() {
	for {
		select {
		case conn := <-h.register:
			h.connections[conn] = true
		case conn := <-h.unregister:
			if _, ok := h.connections[conn]; ok {
				delete(h.connections, conn)
				close(conn.send)
			}
		case message := <-h.broadcast:
			var oMsg models.ActionMessage
			err := json.Unmarshal(message, &oMsg)
			if err != nil {
				log.Printf("Error unmarshalling fired event: %s", string(message[:]))
			} else {
				events.Fire(oMsg.Action, oMsg.Payload)
			}

			/*
				for conn := range h.connections {
					select {
					case conn.send <- message:
					default:
						close(conn.send)
						delete(hub.connections, conn)
					}
				}
			*/
		}
	}
}
