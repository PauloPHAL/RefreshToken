package valueobjects

import (
	"strings"

	"github.com/PauloPHAL/microservices/pkg/interfaces"
	"github.com/PauloPHAL/microservices/pkg/perrors"
)

type Password struct {
	value []byte
}

func NewPassword(password string, passwordManager interfaces.PasswordManager) (*Password, error) {
	password = strings.TrimSpace(password)

	if password == "" {
		return nil, perrors.ErrPasswordRequired
	}

	if len(password) < 6 {
		return nil, perrors.ErrPasswordTooShort
	}

	if len(password) > 128 {
		return nil, perrors.ErrPasswordTooLong
	}

	hashedPassword, err := passwordManager.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &Password{value: hashedPassword}, nil
}

func (p *Password) String() string {
	return "***"
}

func (p *Password) Value() []byte {
	return p.value
}
