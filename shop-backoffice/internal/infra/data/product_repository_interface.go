package data

import (
	"context"

	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/data/model"
)

type ProductRepositoryInterface interface {
	SaveProduct(context.Context, model.ProductModel) error

	GetProductsByName(context.Context, string) ([]model.ProductModel, error)
}
