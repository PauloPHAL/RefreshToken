package security

import (
	"github.com/PauloPHAL/microservices/pkg/interfaces"
	"github.com/PauloPHAL/microservices/pkg/perrors"
	"golang.org/x/crypto/bcrypt"
)

type PasswordManager struct {
	cost int
}

func NewPasswordManager(cost int) interfaces.PasswordManager {
	return &PasswordManager{cost: cost}
}

func (pm *PasswordManager) HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), pm.cost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func (pm *PasswordManager) ComparePasswords(hashedPassword, plainPassword []byte) error {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, plainPassword); err != nil {
		return perrors.ErrInvalidPassword
	}

	return nil
}
