# RapidChat

RapidChat is a real-time chat application. It allows users to join chat rooms, send messages, and interact with other users in real-time. The application is built using React for the frontend and Go for the backend.

## Networking Implementation

RapidChat leverages WebSocket technology for real-time communication between the client and server. WebSocket is a communication protocol that provides **full-duplex** communication channels over a single TCP connection. It enables interactive communication between a client and a server, allowing messages to be sent and received in real-time.

### Frontend

The frontend of RapidChat is built using React. The frontend communicates with the backend server using WebSocket to send and receive messages. The `ChatSocket.js` file handles WebSocket communication with the server, allowing users to join chat rooms, send messages, and receive updates in real-time.

### Backend

The backend of RapidChat is built using Go, a statically typed, compiled programming language. The backend server implements WebSocket functionality using the `github.com/gorilla/websocket` package. It manages WebSocket connections, maintains chat rooms, and facilitates communication between clients.

- **WebSocket Server**: The backend server establishes a WebSocket server using the `github.com/gorilla/websocket` package. It listens for incoming WebSocket connections from clients and manages these connections using a connection pool.

- **Connection Pool**: The backend maintains a connection pool to manage WebSocket connections from multiple clients. When a client connects to the server, its WebSocket connection is added to the connection pool. The server broadcasts messages to all clients in the pool, allowing real-time communication between users.

- **Message Handling**: The backend server handles incoming messages from clients and broadcasts them to other clients in the same chat room. It also sends notifications to clients when users join or leave the chat room.

## Getting Started

To run the RapidChat application locally, follow these steps:

1. Clone the repository.
2. Navigate to the `frontend` directory and run `npm install` to install dependencies.
3. Run `npm start` to start the frontend server.
4. Navigate to the `go-backend` directory and run `go run main.go` to start the backend server.

## Technologies Used

- **Frontend**: React, JavaScript, WebSocket
- **Backend**: Go, WebSocket
- **Package Manager**: npm (for frontend), Go modules (for backend)

## Example

## Acknowledgements

- [Tutorial: Building a Real-Time Chat App with React and Go](https://tutorialedge.net/projects/chat-system-in-go-and-react/)

- [ereid7: go-react-Chat](<https://github.com/ereid7/go-react-chat>)
