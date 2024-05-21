package data

import (
	"context"

	"github.com/hugovallada/shop-poc/shop-backoffice/infra/data/model"
)

type ProductRepositoryInterface interface {
	SaveProduct(context.Context, model.ProductModel) error
}
