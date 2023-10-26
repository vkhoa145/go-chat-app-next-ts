package application

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vkhoa145/go-chat-next-ts/app/server"
	"github.com/vkhoa145/go-chat-next-ts/config"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := config.LoadConfig()
	server := server.NewServer(config)
	err1 := server.Start()
	if err1 != nil {
		log.Fatal("error starting server:", err1)
	}

}
