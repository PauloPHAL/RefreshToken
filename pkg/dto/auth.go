package dto

import "github.com/PauloPHAL/refreshtoken/pkg/perrors"

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (l *LoginDTO) Validate() error {
	if l.Email == "" {
		return perrors.ErrEmailRequired
	}
	if l.Password == "" {
		return perrors.ErrPasswordRequired
	}
	return nil
}

func (r *RefreshTokenDTO) Validate() error {
	if r.RefreshToken == "" {
		return perrors.ErrInvalidToken
	}
	return nil
}
