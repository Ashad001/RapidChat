package websocket

import (
	"fmt"
	"time"
)

// Register - Register channel will send out New User Joined... to all of the clients within this pool when a new client connects.
// Unregister - Will unregister a user and notify the pool when a client disconnects.
// Clients - a map of clients to a boolean value. We can use the boolean value to dictate active/inactive but not disconnected further down the line based on browser focus.
// Broadcast - a channel which, when it is passed a message, will loop through all clients in the pool and send the message through the socket connection.
// _messageList - the messages for the current client
// _messageLimit - the maximum number of messages to send to the socket connection

type Pool struct {
	Register              chan *Client
	Unregister            chan *Client
	Clients               map[*Client]bool
	Rooms                 map[string]map[*Client]bool
	Broadcast             chan Message
	_messageList          []Message
	_messageLimit         int
	_expireAfter_Hrs      time.Duration
	_cleanupMessagesAfter time.Duration
}

func NewPool(messageLimit int, expireAfter_Hrs time.Duration, cleanupMessagesAfter time.Duration) *Pool {
	return &Pool{
		Register:              make(chan *Client),
		Unregister:            make(chan *Client),
		Clients:               make(map[*Client]bool),
		Broadcast:             make(chan Message),
		Rooms:                 make(map[string]map[*Client]bool),
		_messageList:          []Message{},
		_messageLimit:         messageLimit,
		_expireAfter_Hrs:      expireAfter_Hrs,
		_cleanupMessagesAfter: cleanupMessagesAfter,
	}
}

func (pool *Pool) GetUserNames(room string) []UserInfo {
	clients := make([]UserInfo, 0)
	for client := range pool.Rooms[room] {
		clients = append(clients, UserInfo{
			Name:  string(client.User),
			Color: client.Color,
		})
	}
	return clients
}

func (pool *Pool) CleanUpHeartBeat() {
	for range time.Tick(time.Minute * pool._cleanupMessagesAfter) {
		pool.CleanUpMessageList()
	}
}

func (pool *Pool) CleanUpMessageList() {
	if len(pool._messageList) > pool._messageLimit {
		// pool._messageList = pool._messageList[:pool._messageLimit]
		pool._messageList = pool._messageList[len(pool._messageList)-pool._messageLimit:]

	}
	for index, message := range pool._messageList {
		expirationTime := time.Now().Add(-pool._expireAfter_Hrs * time.Hour)
		messageTime, _ := time.Parse(time.RFC3339, message.TimeStamp)
		if messageTime.Before(expirationTime) {
			pool._messageList = pool._messageList[len(pool._messageList)-index:]
			return
		}
	}
}

func (pool *Pool) JoinRoom(client *Client, roomName string) {
	if client.Room != "" {
		delete(pool.Rooms[client.Room], client)
	}

	if pool.Rooms[roomName] == nil {
		pool.Rooms[roomName] = make(map[*Client]bool)
	}

	pool.Rooms[roomName][client] = true
	client.Room = roomName
}

func (pool *Pool) leaveRoom(client *Client) {
	delete(pool.Rooms[client.Room], client)
	if len(pool.Rooms[client.Room]) == 0 {
		delete(pool.Rooms, client.Room)
	}
}

func (pool *Pool) GetClientsInRoom(roomName string) map[*Client]bool {
	return pool.Rooms[roomName]
}

func (pool *Pool) Start() {
	go pool.CleanUpHeartBeat()
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			if pool.Rooms[client.Room] == nil {
				pool.Rooms[client.Room] = make(map[*Client]bool)
			}
			pool.Rooms[client.Room][client] = true
			newUser := string(client.User)
			fmt.Println("Size of Connection Pool (after adding): ", len(pool.Clients))
			clients := pool.Rooms[client.Room]
			for client := range clients {
				client.Conn.WriteJSON(
					Message{
						Type:      1,
						Body:      newUser + " just joined the party!!!",
						TimeStamp: time.Now().Format(time.RFC3339),
					},
				)
				client.Conn.WriteJSON(
					StateMessage{
						Type:       0,
						ClientList: pool.GetUserNames(client.Room),
					},
				)
				pool.CleanUpMessageList()
			}
			fmt.Println(pool.GetUserNames(client.Room))
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			pool.leaveRoom(client)
			userGone := string(client.User)
			fmt.Println("Size of Connection Pool (after deleting): ", len(pool.Clients))
			clients := pool.Rooms[client.Room]
			for client := range clients {
				client.Conn.WriteJSON(
					Message{
						Type:      1,
						Body:      userGone + " left the chat!",
						TimeStamp: time.Now().Format(time.RFC3339),
					},
				)
				client.Conn.WriteJSON(
					StateMessage{
						Type:       0,
						ClientList: pool.GetUserNames(client.Room),
					},
				)
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in a Pool")
			clients := pool.Rooms[message.Room]
			for client := range clients {
				pool.CleanUpMessageList()
				pool._messageList = append(pool._messageList, message)
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					break
				}
			}
		}
	}
}
