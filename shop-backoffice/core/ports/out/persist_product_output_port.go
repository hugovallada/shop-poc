package out

import "github.com/hugovallada/shop-poc/shop-backoffice/core/ports/out/dto"

type PersistProductOutputPort interface {
	Execute(dto.PersistProductParameter)
}