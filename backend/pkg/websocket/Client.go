package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)


type Client struct {
	ID 			string
	UserName 	string
	Color 		string
	Conn 		*websocket.Conn
	Pool 		*Pool
}

type Message struct {
    Type 		int    	`json:"type"`
    Body 		string 	`json:"body"`
	UserName 	string 	`json:"userName"`
	Color 		string	`json:"color"`
	TimeStamp	string 	`json:"timeStamp"`

}

type MessageData struct {
	Message string
	ID 		string
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()
	
	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		var messageData MessageData
		json.Unmarshal([]byte(p), &messageData) 
		if messageData.ID != c.ID {
			log.Println("Unauthorized User")
			return
		}
		message := Message {
			Type: messageType,
			Body: messageData.Message,
			UserName: c.UserName,
			Color: c.Color,
			TimeStamp: time.Now().Format(time.RFC850) }

		c.Pool.Broadcast <- message
		fmt.Printf("Message Recieved %+v\n", message)
	}
}