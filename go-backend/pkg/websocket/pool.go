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
	Register 	chan *Client
	Unregister 	chan *Client
	Clients 	map[*Client]bool
	Broadcast 	chan Message
	_messageList []Message
	_messageLimit int
	_expireAfter_Hrs time.Duration
	_cleanupMessagesAfter time.Duration
}

func NewPool(messageLimit int, expireAfter_Hrs time.Duration, cleanupMessagesAfter time.Duration) *Pool {
	return &Pool{
		Register:     make(chan *Client),
		Unregister:   make(chan *Client),
		Clients:      make(map[*Client]bool),
		Broadcast:    make(chan Message),
		_messageList: []Message{},
		_messageLimit: messageLimit,
		_expireAfter_Hrs: expireAfter_Hrs,
		_cleanupMessagesAfter: cleanupMessagesAfter,
	}
}

func (pool *Pool) GetUserNames() []UserInfo {
	clients := make([]UserInfo, len(pool.Clients))
	i := 0
	for k := range pool.Clients {
		clients[i] = UserInfo {
			Name: k.User,
			Color: k.Color,
		}
		i++
	}
	return clients;
}

func (pool *Pool) CleanUpHeartBeat() {
	for range time.Tick(time.Minute * pool._cleanupMessagesAfter) {
		pool.CleanUpMessageList()
	}
}

func (pool *Pool) CleanUpMessageList() {
	if (len(pool._messageList) > pool._messageLimit) {
		// pool._messageList = pool._messageList[:pool._messageLimit]
		pool._messageList = pool._messageList[len(pool._messageList) - pool._messageLimit:]

	}
	for index, message := range pool._messageList {
		expirationTime := time.Now().Add(-pool._expireAfter_Hrs * time.Hour);
		messageTime, _ := time.Parse(time.RFC3339, message.TimeStamp)
		if (messageTime.Before(expirationTime)) {
			pool._messageList = pool._messageList[len(pool._messageList)-index:]
			return
		}
	}
}
 
func (pool *Pool) Start() {
	go pool.CleanUpHeartBeat()
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			newUser := string(client.User)
			fmt.Println("Size of Connection Pool (after adding): ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(
					Message{
						Type: 1, 
						Body: newUser + " just joined the party!!!", 
						TimeStamp: time.Now().Format(time.RFC3339),
					},
				)
				client.Conn.WriteJSON(
					StateMessage{
						Type: 0, 
						ClientList: pool.GetUserNames(),
					},
				)
				pool.CleanUpMessageList()
			}
			fmt.Println(pool.GetUserNames());
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			userGone := string(client.User)
			fmt.Println("Size of Connection Pool (after deleting): ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(
					Message{
						Type: 1, 
						Body: userGone + " left the chat!",
						TimeStamp: time.Now().Format(time.RFC3339),
					},
				)
				client.Conn.WriteJSON(
					StateMessage{
						Type: 0,
						ClientList: pool.GetUserNames(),
					},
				)
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in a Pool")
				
			for client := range pool.Clients {
				pool.CleanUpMessageList();
				pool._messageList = append(pool._messageList, message)
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					break
				}
			}
		}
	}
}