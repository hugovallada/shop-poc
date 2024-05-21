package translators

import (
	"github.com/hugovallada/shop-poc/shop-backoffice/core/entity"
	outputDto "github.com/hugovallada/shop-poc/shop-backoffice/core/ports/out/dto"
)

func FromProductEntityToPersistProductParameter(product entity.Product) outputDto.PersistProductParameter {
	return outputDto.NewCreateProductParameter(
		product.ID.Value, product.Name.Value, product.Department.Value,
		product.Tags.Value, product.Price.Value, product.TotalPrice.Value,
		product.CreatedAt.Value, product.UpdatedAt.Value, product.Quantity,
		product.Active,
	)
}
