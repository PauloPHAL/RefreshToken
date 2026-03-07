package container

import (
	"github.com/PauloPHAL/microservices/internal/handlers"
	"github.com/PauloPHAL/microservices/internal/services"
	"github.com/PauloPHAL/microservices/pkg/interfaces"
	"github.com/PauloPHAL/microservices/pkg/repository"
	"github.com/PauloPHAL/microservices/pkg/security"
	"gorm.io/gorm"
)

type Container struct {
	UserRepository interfaces.UserRepository
	AuthRepository interfaces.AuthRepository

	UserService interfaces.UserService
	AuthService interfaces.AuthService

	TokenGenerator  interfaces.TokenGenerator
	PasswordManager interfaces.PasswordManager

	UserHandler *handlers.UserHandler
	AuthHandler *handlers.AuthHandler
}

func NewContainer(database *gorm.DB, jwtSecret string, passwordCost int) *Container {
	tokenGenerator := security.NewTokenGenerator(jwtSecret)
	passwordManager := security.NewPasswordManager(passwordCost)

	userRepo := repository.NewUserRepository(database)
	authRepo := repository.NewAuthRepository(database)

	userService := services.NewUserService(userRepo, passwordManager)
	authService := services.NewAuthService(authRepo, tokenGenerator, passwordManager)

	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	return &Container{
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}
}
