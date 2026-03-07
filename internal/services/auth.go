package services

import (
	"context"
	"time"

	"github.com/PauloPHAL/refreshtoken/pkg/dto"
	"github.com/PauloPHAL/refreshtoken/pkg/interfaces"
	"github.com/PauloPHAL/refreshtoken/pkg/models"
	"github.com/PauloPHAL/refreshtoken/pkg/perrors"
)

type AuthServiceImpl struct {
	authRepo        interfaces.AuthRepository
	tokenGenerator  interfaces.TokenGenerator
	passwordManager interfaces.PasswordManager
}

// NewAuthService cria uma nova instância do serviço de autenticação
func NewAuthService(
	authRepo interfaces.AuthRepository,
	tokenGenerator interfaces.TokenGenerator,
	passwordManager interfaces.PasswordManager,
) interfaces.AuthService {
	return &AuthServiceImpl{
		authRepo:        authRepo,
		tokenGenerator:  tokenGenerator,
		passwordManager: passwordManager,
	}
}

func (s *AuthServiceImpl) Login(ctx context.Context, auth *dto.LoginDTO) (*dto.LoginResponseDTO, error) {
	// 1. Busca o usuário (Já garante que user != nil pelo repo)
	user, err := s.authRepo.FindByEmail(ctx, auth.Email)
	if err != nil {
		return nil, perrors.ErrInvalidCredentials
	}

	// 2. Valida a senha
	if err := s.passwordManager.ComparePasswords(user.GetPassword(), []byte(auth.Password)); err != nil {
		return nil, perrors.ErrInvalidCredentials
	}

	var refreshTokenStr string
	refreshTokenEntity := user.GetRefreshToken()

	// 3. Lógica do Refresh Token
	if refreshTokenEntity != nil && time.Now().Before(refreshTokenEntity.GetExpiresAt()) {
		// Reutiliza o token atual se ainda for válido
		refreshTokenStr = refreshTokenEntity.GetToken()
	} else {
		// Se for nil ou expirado, gera um novo
		newRTStr, err := s.tokenGenerator.GenerateRefreshToken(user.GetID())
		if err != nil {
			return nil, err
		}

		// Se a entidade for nil (primeiro login), instanciamos uma nova
		if refreshTokenEntity == nil {
			refreshTokenEntity = &models.RefreshToken{}
		}

		refreshTokenEntity.SetToken(newRTStr)
		refreshTokenEntity.SetUserID(user.GetID())
		refreshTokenEntity.SetExpiresAt(time.Now().Add(time.Hour * 24))

		// Persiste no banco de dados através do repository
		if err := s.authRepo.SaveRefreshToken(ctx, refreshTokenEntity); err != nil {
			return nil, err
		}

		refreshTokenStr = newRTStr
	}

	accessToken, err := s.tokenGenerator.GenerateAccessToken(user.GetID())
	if err != nil {
		return nil, err
	}

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

	// 2. Busca o token específico no banco (Permite múltiplos dispositivos)
	// Se não encontrar, o erro deve ser algo como "sessão não encontrada ou revogada"
	refreshTokenStored, err := s.authRepo.FindRefreshTokenByToken(ctx, refreshTokenDTO.RefreshToken)
	if err != nil {
		return nil, perrors.ErrInvalidToken
	}

	if refreshTokenStored.GetUserID() != userID {
		return nil, perrors.ErrInvalidToken
	}

	if refreshTokenStored.IsExpired() {
		_ = s.authRepo.InvalidateRefreshToken(ctx, refreshTokenDTO.RefreshToken)
		return nil, perrors.ErrTokenExpired
	}

	accessToken, err := s.tokenGenerator.GenerateAccessToken(userID)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenDTO.RefreshToken,
		ExpiresIn:    15 * 60,
	}, nil
}

// ValidateToken valida um token
func (s *AuthServiceImpl) ValidateToken(ctx context.Context, token string) (string, error) {
	userID, err := s.tokenGenerator.ValidateToken(token, "access")
	if err != nil {
		return "", perrors.ErrInvalidToken
	}

	return userID, nil
}

// Logout invalida o refresh token do usuário
func (s *AuthServiceImpl) Logout(ctx context.Context, userID string) error {
	return s.authRepo.InvalidateRefreshToken(ctx, userID)
}
