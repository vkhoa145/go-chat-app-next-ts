package ws

import (
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
