package valueobjects

import (
	"github.com/google/uuid"
)

type ID struct {
	value string
}

func NewID() *ID {
	return &ID{value: uuid.New().String()}
}

func (id *ID) Value() string {
	return id.value
}
