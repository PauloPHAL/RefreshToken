package dto

import "github.com/PauloPHAL/microservices/pkg/perrors"

type UserDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponseDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *UserDTO) Validate() error {
	if u.Name == "" {
		return perrors.ErrNameRequired
	}
	if u.Email == "" {
		return perrors.ErrEmailRequired
	}
	if u.Password == "" {
		return perrors.ErrPasswordRequired
	}
	return nil
}
