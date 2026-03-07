package interfaces

import (
	"context"

	"github.com/PauloPHAL/microservices/pkg/dto"
)

type UserService interface {
	CreateUser(ctx context.Context, user *dto.UserDTO) (*dto.UserResponseDTO, error)
	GetUser(ctx context.Context, id string) (*dto.UserResponseDTO, error)
}

type AuthService interface {
	Login(ctx context.Context, auth *dto.LoginDTO) (*dto.LoginResponseDTO, error)
	Refresh(ctx context.Context, refreshTokenDTO *dto.RefreshTokenDTO) (*dto.LoginResponseDTO, error)
	ValidateToken(ctx context.Context, token string) (string, error)
	Logout(ctx context.Context, userID string) error
}
