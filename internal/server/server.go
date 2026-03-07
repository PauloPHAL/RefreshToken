package server

import (
	"github.com/PauloPHAL/microservices/internal/config"
	"github.com/PauloPHAL/microservices/internal/container"
	"github.com/PauloPHAL/microservices/internal/router"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Start(database *gorm.DB) {
	rotas := gin.Default()

	cfg := config.GetConfig()
	container := container.NewContainer(database, cfg.GetJWTSecret(), cfg.GetPasswordCost())

	router.Api(rotas, container)

	rotas.Run()
}
