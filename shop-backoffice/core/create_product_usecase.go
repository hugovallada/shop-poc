package core

import (
	"github.com/hugovallada/shop-poc/shop-backoffice/core/entity/factorys"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/ports/in/dto"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/ports/out"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/translators"
)

type CreateProductUseCase struct {
	persistProductOutputPort out.PersistProductOutputPort
}

func (cp CreateProductUseCase) Execute(createProductParameter dto.CreateProductParameter) {
	product := factorys.ProductFactoryFromCreateProductParameter(createProductParameter)
	product.UpdateProduct()
	persistProductParameter := translators.FromProductEntityToPersistProductParameter(product)
	cp.persistProductOutputPort.Execute(persistProductParameter)
}
