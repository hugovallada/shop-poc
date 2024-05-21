package vo

import (
	"github.com/google/uuid"
)

type ID struct {
	Value uuid.UUID
}

func NewID() ID {
	return ID{Value: uuid.New()}
}
