package websocket

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	ID		string
	Room 	string
	User	string
	Color	string
	Conn	*websocket.Conn
	Pool	*Pool
}

type UserInfo struct {
	Name string `json:"name"`
	Color string `json:"color"`
}

type Message struct {
    Type 		int    	`json:"type"`
    Body 		string 	`json:"body"`
	User 		string 	`json:"user"`
	Color 		string	`json:"color"`
	Room 		string 	`json:"room"`
	TimeStamp	string 	`json:"timeStamp"`
}

type MessageData struct {
	Message string
	Id 		string
}


type StateMessage struct {
	Type int `json:"type"`
	ClientList []UserInfo `json:"clientList"`
}