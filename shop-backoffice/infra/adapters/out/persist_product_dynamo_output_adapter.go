package out

import (
	"context"
	"log/slog"

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

func (p *PersistProductOutputAdapter) Execute(ctx context.Context, persistProductParameter dto.PersistProductParameter) error {
	productModel := model.NewProductModel(persistProductParameter)
	slog.InfoContext(ctx,
		"new ProductModel created from PersistProductParamter",
		slog.Any("persistProductParamter", persistProductParameter), slog.Any("productModel", productModel))
	return p.productRepository.SaveProduct(ctx, productModel)
}
