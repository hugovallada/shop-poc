package out

import (
	"context"

	"github.com/hugovallada/shop-poc/shop-backoffice/internal/core/ports/out/dto"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/data"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/translators"
)

type GetProductByNameOutputAdapter struct {
	repository data.ProductRepositoryInterface
}

func NewGetProductByNameOutputAdapter(repository data.ProductRepositoryInterface) GetProductByNameOutputAdapter {
	return GetProductByNameOutputAdapter{repository: repository}
}

func (gn GetProductByNameOutputAdapter) Execute(ctx context.Context, name string) ([]dto.GetProductByNameResponse, error) {
	models, err := gn.repository.GetProductsByName(ctx, name)
	if err != nil {
		return nil, err
	}
	var productsResponse []dto.GetProductByNameResponse
	for _, model := range models {
		productsResponse = append(productsResponse, translators.FromProductModelToGetProductByNameResponseImpl(model))
	}
	return productsResponse, nil
}
