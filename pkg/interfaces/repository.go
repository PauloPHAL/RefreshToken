package interfaces

import (
	"context"

	"github.com/PauloPHAL/refreshtoken/pkg/dto"
	"github.com/PauloPHAL/refreshtoken/pkg/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *dto.UserDTO, passwordManager PasswordManager) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	UserExists(ctx context.Context, email string) (bool, error)
}

type AuthRepository interface {
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	SaveRefreshToken(ctx context.Context, token *models.RefreshToken) error
	FindRefreshTokenByToken(ctx context.Context, token string) (*models.RefreshToken, error)
	InvalidateRefreshToken(ctx context.Context, userID string) error
}

type TokenGenerator interface {
	GenerateAccessToken(userID string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
	ValidateToken(tokenString, expectedType string) (string, error)
}

type PasswordManager interface {
	HashPassword(password string) ([]byte, error)
	ComparePasswords(hashedPassword, plainPassword []byte) error
}
