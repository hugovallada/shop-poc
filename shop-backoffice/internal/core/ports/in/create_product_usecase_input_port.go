package in

import (
	"context"

	"github.com/hugovallada/shop-poc/shop-backoffice/internal/core/ports/in/dto"
)

type CreateProductUseCaseInputPort interface {
	Execute(context.Context, dto.CreateProductParameter) error
}
