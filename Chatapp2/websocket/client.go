package ws

import (
	store "chatapp2/connectdb"
	"chatapp2/model"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn // actual WebSocket connection
	Send chan OutboundMessage
	Hub  *Hub
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		var msg InboundMessage
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			break
		}
		store.DB.Create(&model.Message{
			SenderID:   msg.SenderID,
			ReceiverID: msg.ReceiverID,
			Content:    msg.Content,
		})
		c.Hub.BroadCast <- OutboundMessage{
			SenderID:   msg.SenderID,
			ReceiverID: msg.ReceiverID,
			Content:    msg.Content,
		}

	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()
	for msg := range c.Send {
		err := c.Conn.WriteJSON(msg)
		if err != nil {
			break
		}
	}
}
