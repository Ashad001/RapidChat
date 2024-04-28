package main

import (
	"fmt"
	"net/http"

	"github.com/Ashad001/RapidChat/pkg/utils"
	"github.com/Ashad001/RapidChat/pkg/websocket"
)

type ChatServer struct {
	messageList []websocket.MessageData
}



func (c *ChatServer) serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Host)
	fmt.Println("WebSocket Endpoint Hit")

	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	keys := r.URL.Query()
	
	userName := keys.Get("user")
	if len(userName) < 1 {
		fmt.Println("URl Parameter 'user' is missing")
		return
	}
	
	userId := keys.Get("userId")
	if len(userId) < 1 {
		fmt.Println("Url Parameter 'userId' is missing")
		return
	}

	color := utils.GetRandomColor()
	fmt.Println(color)

	client := &websocket.Client{
		ID: userId,
		User: userName,
		Color: color,
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func (c *ChatServer) setupRoutes() {
	fmt.Println("Distributed Chat Server")
	pool := websocket.NewPool(10, 20, 30)
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		c.serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Chat App")
	chatServer := ChatServer{make([]websocket.MessageData, 0)}
	chatServer.setupRoutes()
	
	http.ListenAndServe(":8080", nil)
}

