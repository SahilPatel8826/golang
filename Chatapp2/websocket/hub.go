package ws

type Hub struct {
	Clients    map[*Client]bool
	BroadCast  chan OutboundMessage
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		BroadCast:  make(chan OutboundMessage),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case msg := <-h.BroadCast:
			for client := range h.Clients {
				if client.ID == msg.ReceiverID || client.ID == msg.SenderID {
					client.Send <- msg
				}
			}
		}
	}
}
