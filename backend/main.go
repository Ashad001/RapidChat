package main

import (
	"fmt"
	"net/http"

	"github.com/Ashad001/RapidChat/pkg/websocket"
	
)



func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
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

