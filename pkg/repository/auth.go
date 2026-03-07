package repository

import (
	"context"

	"github.com/PauloPHAL/microservices/pkg/models"
	"github.com/PauloPHAL/microservices/pkg/perrors"
	"github.com/PauloPHAL/microservices/pkg/valueobjects"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	database *gorm.DB
}

func NewAuthRepository(database *gorm.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{database: database}
}

func (r *AuthRepositoryImpl) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.database.WithContext(ctx).Preload("RefreshToken").Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, perrors.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepositoryImpl) SaveRefreshToken(ctx context.Context, token *models.RefreshToken) error {
	token.SetID(valueobjects.NewID().Value())

	r.database.WithContext(ctx).Delete(&models.RefreshToken{}, "user_id = ?", token.GetUserID())

	result := r.database.WithContext(ctx).Create(&models.RefreshToken{
		ID:        token.GetID(),
		Token:     token.GetToken(),
		UserID:    token.GetUserID(),
		ExpiresAt: token.GetExpiresAt(),
	})

	return result.Error
}

func (r *AuthRepositoryImpl) FindRefreshTokenByToken(ctx context.Context, token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	if err := r.database.WithContext(ctx).Where("token = ?", token).First(&refreshToken).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &refreshToken, nil
}

func (r *AuthRepositoryImpl) InvalidateRefreshToken(ctx context.Context, userID string) error {
	return r.database.WithContext(ctx).Delete(&models.RefreshToken{}, "user_id = ?", userID).Error
}
