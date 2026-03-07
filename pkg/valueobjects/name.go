package valueobjects

import (
	"strings"

	"github.com/PauloPHAL/microservices/pkg/perrors"
)

type Name struct {
	value string
}

func NewName(name string) (*Name, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, perrors.ErrNameRequired
	}

	if len(name) < 2 {
		return nil, perrors.ErrNameTooShort
	}

	if len(name) > 100 {
		return nil, perrors.ErrNameTooLong
	}

	return &Name{value: name}, nil
}

func (n *Name) String() string {
	return n.value
}

func (n *Name) Value() string {
	return n.value
}

func (n *Name) Equals(other *Name) bool {
	if other == nil {
		return false
	}
	return strings.EqualFold(n.value, other.value)
}
