package mocks

import (
	"context"
	"errors"
	"strings"

	"github.com/hugovallada/shop-poc/shop-backoffice/core/ports/out/dto"
)

type PersistProductOutputPortMock struct{}

func (p *PersistProductOutputPortMock) Execute(ctx context.Context, p0 dto.PersistProductParameter) error {
	if strings.HasPrefix(p0.GetName(), "FAILURE") {
		return errors.New("invalid product")
	}
	return nil
}
