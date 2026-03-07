package valueobjects

import (
	"regexp"
	"strings"

	"github.com/PauloPHAL/microservices/pkg/perrors"
)

type Email struct {
	value string
}

func NewEmail(email string) (*Email, error) {
	email = strings.TrimSpace(email)

	if email == "" {
		return nil, perrors.ErrEmailRequired
	}

	const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(emailRegex)

	if !regex.MatchString(email) {
		return nil, perrors.ErrInvalidEmail
	}

	if len(email) > 254 {
		return nil, perrors.ErrEmailTooLong
	}

	return &Email{value: email}, nil
}

func (e *Email) String() string {
	return e.value
}

func (e *Email) Value() string {
	return e.value
}

func (e *Email) Equals(other *Email) bool {
	if other == nil {
		return false
	}
	return strings.EqualFold(e.value, other.value)
}
