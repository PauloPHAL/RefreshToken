package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/PauloPHAL/refreshtoken/internal/config"
	"github.com/PauloPHAL/refreshtoken/pkg/dto"
	"github.com/PauloPHAL/refreshtoken/pkg/interfaces"
	"github.com/PauloPHAL/refreshtoken/pkg/models"
	"github.com/PauloPHAL/refreshtoken/pkg/perrors"
)

type AuthServiceImpl struct {
	authRepo        interfaces.AuthRepository
	tokenGenerator  interfaces.TokenGenerator
	passwordManager interfaces.PasswordManager
	cache           *config.Cache
}

func NewAuthService(
	authRepo interfaces.AuthRepository,
	tokenGenerator interfaces.TokenGenerator,
	passwordManager interfaces.PasswordManager,
	cache *config.Cache,
) interfaces.AuthService {
	return &AuthServiceImpl{
		authRepo:        authRepo,
		tokenGenerator:  tokenGenerator,
		passwordManager: passwordManager,
		cache:           cache,
	}
}

func (s *AuthServiceImpl) Login(ctx context.Context, auth *dto.LoginDTO) (*dto.LoginResponseDTO, error) {
	user, err := s.authRepo.FindByEmail(ctx, auth.Email)
	if err != nil {
		return nil, perrors.ErrInvalidCredentials
	}

	if err := s.passwordManager.ComparePasswords(user.GetPassword(), []byte(auth.Password)); err != nil {
		return nil, perrors.ErrInvalidCredentials
	}

	var refreshTokenStr string
	refreshTokenEntity := user.GetRefreshToken()

	if refreshTokenEntity != nil && time.Now().Before(refreshTokenEntity.GetExpiresAt()) {
		refreshTokenStr = refreshTokenEntity.GetToken()
	} else {
		newRTStr, err := s.tokenGenerator.GenerateRefreshToken(user.GetID())
		if err != nil {
			return nil, err
		}

		if refreshTokenEntity == nil {
			refreshTokenEntity = &models.RefreshToken{}
		}

		refreshTokenEntity.SetToken(newRTStr)
		refreshTokenEntity.SetUserID(user.GetID())
		refreshTokenEntity.SetExpiresAt(time.Now().Add(time.Hour * 24))

		if err := s.authRepo.SaveRefreshToken(ctx, refreshTokenEntity); err != nil {
			return nil, err
		}

		refreshTokenStr = newRTStr
	}

	accessToken, err := s.tokenGenerator.GenerateAccessToken(user.GetID())
	if err != nil {
		return nil, err
	}

	s.cache.Set("userRefresh:"+user.GetID(), refreshTokenEntity, time.Hour*24)
	s.cache.Set("userAccess:"+user.GetID(), accessToken, time.Minute*15)

	return &dto.LoginResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenStr,
		ExpiresIn:    15 * 60,
	}, nil
}

func (s *AuthServiceImpl) Refresh(ctx context.Context, refreshTokenDTO *dto.RefreshTokenDTO) (*dto.LoginResponseDTO, error) {
	userID, err := s.tokenGenerator.ValidateToken(refreshTokenDTO.RefreshToken, "refresh")
	if err != nil {
		return nil, err
	}

	accessToken, err := s.tokenGenerator.GenerateAccessToken(userID)
	if err != nil {
		return nil, err
	}

	var refreshTokenStored *models.RefreshToken
	cachedData, err := s.cache.Get("userRefresh:" + userID)
	if err == nil {
		var cachedToken models.RefreshToken
		if err := json.Unmarshal(cachedData, &cachedToken); err == nil {
			refreshTokenStored = &cachedToken
		}
	} else {
		refreshTokenStored, err = s.authRepo.FindRefreshTokenByToken(ctx, refreshTokenDTO.RefreshToken)
		if err != nil {
			return nil, perrors.ErrInvalidToken
		}
		s.cache.Set("userRefresh:"+userID, refreshTokenStored, time.Hour*24)
		s.cache.Set("userAccess:"+userID, accessToken, time.Minute*15)
	}

	if refreshTokenStored.GetUserID() != userID {
		s.cache.Delete("userRefresh:" + userID)
		s.cache.Delete("userAccess:" + userID)
		return nil, perrors.ErrInvalidToken
	}

	if refreshTokenStored.IsExpired() {
		s.authRepo.InvalidateRefreshToken(ctx, refreshTokenDTO.RefreshToken)
		s.cache.Delete("userRefresh:" + userID)
		s.cache.Delete("userAccess:" + userID)
		return nil, perrors.ErrTokenExpired
	}

	return &dto.LoginResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenDTO.RefreshToken,
		ExpiresIn:    15 * 60,
	}, nil
}

func (s *AuthServiceImpl) ValidateToken(ctx context.Context, token string) (string, error) {
	userID, err := s.tokenGenerator.ValidateToken(token, "access")
	if err != nil {
		return "", perrors.ErrInvalidToken
	}

	return userID, nil
}

func (s *AuthServiceImpl) Logout(ctx context.Context, userID string) error {
	err := s.authRepo.InvalidateRefreshToken(ctx, userID)
	if err != nil {
		return err
	}
	s.cache.Delete("userRefresh:" + userID)
	s.cache.Delete("userAccess:" + userID)
	return nil
}
