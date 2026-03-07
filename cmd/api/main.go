package main

import (
	"github.com/PauloPHAL/microservices/internal/config"
	"github.com/PauloPHAL/microservices/internal/server"
)

func initialize() {
	config.GetConfig()
	config.GetDB()
	config.SyncDB()
}

func main() {
	initialize()

	server.Start(config.DB)
}
