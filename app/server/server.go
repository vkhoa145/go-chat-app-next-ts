package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vkhoa145/go-chat-next-ts/config"
	"github.com/vkhoa145/go-chat-next-ts/db"
	"gorm.io/gorm"
)

type Server struct {
	Fiber  *fiber.App
	DB     *gorm.DB
	Config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Fiber:  fiber.New(),
		Config: cfg,
		DB:     db.Init(cfg),
	}
}

func (server *Server) Start() error {
	server.Fiber.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type.Accept,Content-Lenght,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowOriginsFunc: func(origin string) bool {
			return origin == "https://localhost:3000"
		},
	}))
	SetupRoutes(server)
	return server.Fiber.Listen(":" + server.Config.HTTP.Port)
}
