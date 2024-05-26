package translators

import (
	"github.com/google/uuid"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/out/dto"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/data/model"
)

func FromProductModelToGetProductByNameResponseImpl(productModel model.ProductModel) dto.GetProductByNameResponseImpl {
	return dto.GetProductByNameResponseImpl{
		ID:         uuid.MustParse(productModel.ID),
		Name:       productModel.Name,
		Department: productModel.Department,
		Tags:       productModel.Tags,
		Price:      productModel.Price,
		Quantity:   productModel.Quantity,
		TotalPrice: productModel.TotalPrice,
		Active:     productModel.Active,
		CreatedAt:  productModel.CreatedAt,
		UpdatedAt:  productModel.UpdatedAt,
	}
}
