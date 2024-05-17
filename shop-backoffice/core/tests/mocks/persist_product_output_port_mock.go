package mocks

import (
	"errors"
	"strings"

	"github.com/hugovallada/shop-poc/shop-backoffice/core/ports/out/dto"
)

type PersistProductOutputPortMock struct{}

func (p *PersistProductOutputPortMock) Execute(p0 dto.PersistProductParameter) error {
	if strings.HasPrefix(p0.GetName(), "FAILURE") {
		return errors.New("invalid product")
	}
	return nil
}
