package out

import (
	"context"

	"github.com/hugovallada/shop-poc/shop-backoffice/core/ports/out/dto"
)

type PersistProductOutputPort interface {
	Execute(context.Context, dto.PersistProductParameter) error
}
