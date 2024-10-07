package main

import (
	"log"
	"video-microservice/pkg/config"
	"video-microservice/pkg/di"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config", err.Error())
	}
	server, err1 := di.InitializeServe(c)
	if err1 != nil {
		log.Fatal("failed to init server", err1.Error())
	}
	if err := server.Start(); err != nil {
		log.Fatal("couldn't start the server")
	}
}