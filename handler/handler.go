package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/grupo8hackglobo2019/back-end/model"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan model.Message)

type Handler struct {
	Upgrader websocket.Upgrader
}

func (h *Handler) HandleConnections(w http.ResponseWriter, r *http.Request) {
	
	ws, err := h.Upgrader.Upgrade(w, r, nil)
	if err != nil {
			log.Fatal(err)
	}
	
	defer ws.Close()

	clients[ws] = true

	for {
		var msg model.Message
		
		err := ws.ReadJSON(&msg)
		if err != nil {
				log.Printf("error: %v", err)
				delete(clients, ws)
				break
		}

		log.Printf("payload reading: %v", msg)
		
		broadcast <- msg
	}

}

func (h *Handler) HandleMessages() {
	for {
			
		msg := <- broadcast
		
		for client := range clients {
				err := client.WriteJSON(msg)
				log.Printf("payload writting: %v", msg)
				if err != nil {
						log.Printf("error: %v", err)
						client.Close()
						delete(clients, client)
				}
		}
	}
}