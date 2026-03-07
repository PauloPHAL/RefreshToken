package repository

import (
	"context"

	"github.com/PauloPHAL/refreshtoken/pkg/dto"
	"github.com/PauloPHAL/refreshtoken/pkg/interfaces"
	"github.com/PauloPHAL/refreshtoken/pkg/models"
	"github.com/PauloPHAL/refreshtoken/pkg/perrors"
	"github.com/PauloPHAL/refreshtoken/pkg/valueobjects"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{database: database}
}

func (u *UserRepositoryImpl) CreateUser(ctx context.Context, userDTO *dto.UserDTO, passwordManager interfaces.PasswordManager) (*models.User, error) {
	var user models.User
	email, err := valueobjects.NewEmail(userDTO.Email)
	if err != nil {
		return nil, err
	}

	exists, err := u.UserExists(ctx, email.Value())
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, perrors.ErrEmailAlreadyExists
	}

	name, err := valueobjects.NewName(userDTO.Name)
	if err != nil {
		return nil, err
	}

	password, err := valueobjects.NewPassword(userDTO.Password, passwordManager)
	if err != nil {
		return nil, err
	}

	user.SetID(valueobjects.NewID().Value())
	user.SetName(name.Value())
	user.SetEmail(email.Value())
	user.SetPassword(password.Value())

	result := u.database.WithContext(ctx).Create(&models.User{
		ID:       user.GetID(),
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := u.database.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, perrors.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserRepositoryImpl) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	if err := u.database.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, perrors.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserRepositoryImpl) UserExists(ctx context.Context, email string) (bool, error) {
	var count int64
	if err := u.database.WithContext(ctx).Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
