package websocket

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
	Clients    map[*Client]bool
}
NewPool()