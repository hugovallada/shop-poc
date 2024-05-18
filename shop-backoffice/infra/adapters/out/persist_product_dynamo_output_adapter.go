package out

import (
	"github.com/hugovallada/shop-poc/shop-backoffice/core/ports/out/dto"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/data"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/data/model"
)

type PersistProductOutputAdapter struct {
	productRepository data.ProductRepository
}

func NewPersistProductDynamoOutputAdapter(productRepository data.ProductRepository) PersistProductOutputAdapter {
	return PersistProductOutputAdapter{
		productRepository: productRepository,
	}
}

func (p *PersistProductOutputAdapter) Execute(persistProductParameter dto.PersistProductParameter) error {
	productModel := model.NewProductModel(persistProductParameter)
	return p.productRepository.SaveProduct(productModel)
}
