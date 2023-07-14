package main

import (
	"log"

	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/config"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/di"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}

	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatal("failed to initialize server: ", err)
	}

	if err := server.Start(); err != nil {
		log.Fatal("failed to start server: ", err)
	}

}
