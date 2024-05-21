package in

import (
	"context"

	"github.com/hugovallada/shop-poc/shop-backoffice/core/ports/in/dto"
)

type CreateProductUseCaseInputPort interface {
	Execute(context.Context, dto.CreateProductParameter) error
}
