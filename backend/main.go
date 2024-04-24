package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool { return true},
}

func reader(conn *websocket.Conn) {
	for {messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return 
		}
		fmt.Println(string(p))
		if err := conn.WriteMessage(messageType, p); 
		err != nil {
			log.Println(err)
			fmt.Println("Error", err)
			return
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		fmt.Println(err)
	}
	reader(ws)
}

func serupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome")
	})
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App")
	serupRoutes()
	
	http.ListenAndServe(":8080", nil)
}

