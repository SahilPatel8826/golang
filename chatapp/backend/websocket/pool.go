package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
	Clients    map[*Client]bool
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message),
		Clients:    make(map[*Client]bool),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool:", len(pool.Clients))
			for client, _ := range pool.Clients {
				fmt.Println("Client Connected:", client.ID)
				client.Conn.WriteJSON(Message{Type: 1, Body: "Welcome to the WebSocket Chat!"})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool:", len(pool.Clients))
			for client, _ := range pool.Clients {
				fmt.Println("Client Disconnected:", client.ID)
				client.Conn.WriteJSON(Message{Type: 1, Body: "A user has left the chat."})
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
