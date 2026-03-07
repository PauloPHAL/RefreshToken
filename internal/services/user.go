package services

import (
	"context"

	"github.com/PauloPHAL/refreshtoken/pkg/dto"
	"github.com/PauloPHAL/refreshtoken/pkg/interfaces"
)

type UserServiceImpl struct {
	repo     interfaces.UserRepository
	password interfaces.PasswordManager
}

func NewUserService(repo interfaces.UserRepository, password interfaces.PasswordManager) interfaces.UserService {
	return &UserServiceImpl{repo: repo, password: password}
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, user *dto.UserDTO) (*dto.UserResponseDTO, error) {
	userRepo, err := u.repo.CreateUser(ctx, user, u.password)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponseDTO{
		ID:    userRepo.GetID(),
		Name:  userRepo.GetName(),
		Email: userRepo.GetEmail(),
	}, nil
}

func (u *UserServiceImpl) GetUser(ctx context.Context, id string) (*dto.UserResponseDTO, error) {
	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponseDTO{
		ID:    user.GetID(),
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}, nil
}
