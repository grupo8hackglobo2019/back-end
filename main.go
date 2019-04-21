package main

import (
	"log"
    "net/http"

	"github.com/gorilla/websocket"
	"github.com/grupo8hackglobo2019/back-end/handler"
)

func main() {

	chatHandler := &handler.Handler{
		Upgrader: websocket.Upgrader{},
	}

	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", chatHandler.HandleConnections)

	go chatHandler.HandleMessages()

	log.Println("http server started on :8000")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
			log.Fatal("ListenAndServe: ", err)
	}

}