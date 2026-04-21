package main

import (
	"github.com/PauloPHAL/refreshtoken/internal/config"
	"github.com/PauloPHAL/refreshtoken/internal/server"
)

func main() {
	config.GetConfig()
	db := config.GetDB()
	cache := config.NewCache()
	config.SyncDB()

	server.Start(db, cache)
}
