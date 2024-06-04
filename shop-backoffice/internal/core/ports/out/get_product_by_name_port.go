package out

import (
	"context"

	"github.com/hugovallada/shop-poc/shop-backoffice/internal/core/ports/out/dto"
)

type GetProductByNameOutputPort interface {
	Execute(context.Context, string) ([]dto.GetProductByNameResponse, error)
}
