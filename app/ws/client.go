package ws

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	message  chan *Message
	ID       string `json:"id"`
	RoomId   string `json:"room_id"`
	Username string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	RoomId   string `json:"room_id"`
	Username string `json:username`
}

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
		}

		msg := &Message{
			Content:  string(m),
			RoomId:   c.RoomId,
			Username: c.Username,
		}
		log.Println("msg:", msg)

		hub.Broadcast <- msg
	}
}
