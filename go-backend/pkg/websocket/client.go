package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

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
		if messageData.Id != c.ID {
			log.Println("Unauthorized User")
			return
		}
		message := Message{
			Type:      messageType,
			Body:      messageData.Message,
			User:      c.User,
			Color:     c.Color,
			Room:      c.Room,
			TimeStamp: time.Now().Format(time.RFC3339),
		}
		c.Pool.Broadcast <- message
	}
}
