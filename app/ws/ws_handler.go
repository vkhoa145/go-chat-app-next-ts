package ws

import (
	"log"
	"net/http"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateRoom struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom() fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := CreateRoom{}

		if err := c.BodyParser(&payload); err != nil {
			c.Status(http.StatusBadRequest)
		}

		h.hub.Rooms[payload.ID] = &Room{
			ID:      payload.ID,
			Name:    payload.Name,
			Clients: make(map[string]*Client),
		}

		c.Status(http.StatusCreated)
		return c.JSON(&fiber.Map{"status": http.StatusCreated, "error": nil, "room": h.hub.Rooms[payload.ID]})
	}
}

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

func (h *Handler) JoinRoom() fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		defer func() {
			c.Conn.Close()
		}()

		roomId := c.Params("roomId")
		clientId := c.Query("userId")
		username := c.Query("username")

		client := &Client{
			Conn:     c,
			message:  make(chan *Message),
			ID:       clientId,
			RoomId:   roomId,
			Username: username,
		}

		message := &Message{
			Content:  "A new user has joined the room",
			RoomId:   roomId,
			Username: username,
		}

		h.hub.Register <- client
		h.hub.Broadcast <- message

		var (
			mesType int
			mess    []byte
			err     error
		)

		for {
			if mesType, mess, err = c.ReadMessage(); err != nil {
				log.Println("read error:", err)
				break
			}

			log.Printf("recv: %s", mess)
			if err = c.WriteMessage(mesType, mess); err != nil {
				log.Println("write:", err)
				break
			}

			msg := &Message{
				Content:  string(mess),
				RoomId:   client.RoomId,
				Username: client.Username,
			}

			h.hub.Broadcast <- msg
		}
	})
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms() fiber.Handler {
	return func(c *fiber.Ctx) error {
		rooms := make([]RoomRes, 0)
		for _, r := range h.hub.Rooms {
			rooms = append(rooms, RoomRes{
				ID:   r.ID,
				Name: r.Name,
			})
		}
		c.Status(http.StatusOK)
		return c.JSON(&fiber.Map{"status": http.StatusOK, "rooms": rooms, "error": nil})
	}
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var clients []ClientRes
		roomId := c.Params("roomId")

		if _, ok := h.hub.Rooms[roomId]; !ok {
			clients = make([]ClientRes, 0)
			c.Status(http.StatusOK)
			return c.JSON(&fiber.Map{"status": http.StatusOK, "clients": clients, "error": nil})

		}

		for _, c := range h.hub.Rooms[roomId].Clients {
			clients = append(clients, ClientRes{
				ID:       c.ID,
				Username: c.Username,
			})
		}
		c.Status(http.StatusOK)
		return c.JSON(&fiber.Map{"status": http.StatusOK, "clients": clients, "error": nil})
	}
}
