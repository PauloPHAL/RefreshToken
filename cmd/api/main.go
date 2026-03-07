package main

import (
	"github.com/PauloPHAL/refreshtoken/internal/config"
	"github.com/PauloPHAL/refreshtoken/internal/server"
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
