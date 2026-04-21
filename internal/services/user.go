package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/PauloPHAL/refreshtoken/internal/config"
	"github.com/PauloPHAL/refreshtoken/pkg/dto"
	"github.com/PauloPHAL/refreshtoken/pkg/interfaces"
	"github.com/PauloPHAL/refreshtoken/pkg/models"
)

type UserServiceImpl struct {
	repo     interfaces.UserRepository
	password interfaces.PasswordManager
	cache    *config.Cache
}

func NewUserService(repo interfaces.UserRepository, password interfaces.PasswordManager, cache *config.Cache) interfaces.UserService {
	return &UserServiceImpl{repo: repo, password: password, cache: cache}
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
	cachedData, err := u.cache.Get("user:" + id)
	if err == nil {
		var user models.User
		if err := json.Unmarshal(cachedData, &user); err == nil {
			return &dto.UserResponseDTO{
				ID:    user.GetID(),
				Name:  user.GetName(),
				Email: user.GetEmail(),
			}, nil
		}
	}

	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	u.cache.Set("user:"+id, user, time.Hour)

	return &dto.UserResponseDTO{
		ID:    user.GetID(),
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}, nil
}
