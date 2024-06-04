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

func ParseIdOrNew(stringId string) ID {
	var id uuid.UUID
	id, err := uuid.Parse(stringId)
	if err != nil {
		id = uuid.New()
	}
	return ID{Value: id}
}
