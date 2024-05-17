package in

import "github.com/hugovallada/shop-poc/shop-backoffice/core/ports/in/dto"

type CreateProductUseCaseInputPort interface {
	Execute(dto.CreateProductParameter) error
}
