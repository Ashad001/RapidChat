package websocket

import "fmt"

// Register - Register channel will send out New User Joined... to all of the clients within this pool when a new client connects.
// Unregister - Will unregister a user and notify the pool when a client disconnects.
// Clients - a map of clients to a boolean value. We can use the boolean value to dictate active/inactive but not disconnected further down the line based on browser focus.
// Broadcast - a channel which, when it is passed a message, will loop through all clients in the pool and send the message through the socket connection.

type Pool struct {
	Register 	chan *Client
	Unregister 	chan *Client
	Clients 	map[*Client]bool
	Broadcast 	chan Message
}

func NewPool() *Pool {
	return &Pool {
		Register: 	make(chan *Client),
		Unregister: make(chan *Client),
		Clients: 	make(map[*Client]bool),
		Broadcast: 	make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool (after adding): ", len(pool.Clients))
			for client, _ := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New Peep Here"})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool (after deleting): ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected... "})
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in a Pool")
				
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					break
				}
			}
		}
	}
}