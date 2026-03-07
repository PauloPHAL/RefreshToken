package router

import (
	"github.com/PauloPHAL/refreshtoken/internal/container"
	"github.com/gin-gonic/gin"
)

func Api(rotas *gin.Engine, c *container.Container) {

	v1 := rotas.Group("/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("", c.UserHandler.CreateUser)
			user.GET("", c.UserHandler.GetUser)
		}

		auth := v1.Group("/auth")
		{
			auth.POST("/login", c.AuthHandler.Login)
			auth.POST("/refresh", c.AuthHandler.Refresh)
			auth.GET("/validate", c.AuthHandler.AuthJWT, c.AuthHandler.ValidateToken)
			auth.POST("/logout", c.AuthHandler.AuthJWT, c.AuthHandler.Logout)
		}
	}

}
