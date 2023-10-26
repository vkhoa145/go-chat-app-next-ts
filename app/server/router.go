package server

import (
	UserHandler "github.com/vkhoa145/go-chat-next-ts/app/modules/users/handlers"
	UserRepositoy "github.com/vkhoa145/go-chat-next-ts/app/modules/users/repositories"
	UserUsecase "github.com/vkhoa145/go-chat-next-ts/app/modules/users/usecase"
	"github.com/vkhoa145/go-chat-next-ts/app/ws"
)

func SetupRoutes(server *Server) {
	userRepo := UserRepositoy.NewUserRepo(server.DB)
	userUsecase := UserUsecase.NewUserUsecase(userRepo)
	userHandler := UserHandler.NewUserHandler(userRepo, userUsecase)

	api := server.Fiber
	user := api.Group("/users")
	user.Post("/signup", userHandler.SignUp(server.Config))
	user.Post("/signin", userHandler.SignIn(server.Config))

	// api.Use("ws", func(c *fiber.Ctx) error {
	// 	if websocket.IsWebSocketUpgrade(c) {
	// 		c.Locals("allowed", true)
	// 		return c.Next()
	// 	}

	// 	return fiber.ErrUpgradeRequired
	// })

	// api.Get("ws/:id", websocket.New(func(c *websocket.Conn) {
	// 	log.Println(c.Locals("allowed"))

	// 	var (
	// 		mesType int
	// 		mess    []byte
	// 		err     error
	// 	)

	// 	for {
	// 		if mesType, mess, err = c.ReadMessage(); err != nil {
	// 			log.Println("read error:", err)
	// 			break
	// 		}

	// 		log.Printf("recv: %s", mess)
	// 		if err = c.WriteMessage(mesType, mess); err != nil {
	// 			log.Println("write:", err)
	// 			break
	// 		}
	// 	}
	// }))
	hub := ws.NewHub()
	wsRoute := api.Group("/ws")

	wsHandler := ws.NewHandler(hub)
	wsRoute.Post("room", wsHandler.CreateRoom())

	go hub.Run()
	
	wsRoute.Get("/joinRoom/:roomId", wsHandler.JoinRoom())
	wsRoute.Get("/getRooms", wsHandler.GetRooms())
	wsRoute.Get("/getClients/:roomId", wsHandler.GetClients())
}
