package translators

import (
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/core/entity"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/core/ports/out/dto"
)

func FromProductEntityToPersistProductParameter(product entity.Product) dto.PersistProductParameter {
	return dto.NewCreateProductParameter(
		product.ID.Value, product.Name.Value, product.Department.Value,
		product.Tags.Value, product.Price.Value, product.TotalPrice.Value,
		product.CreatedAt.Value, product.UpdatedAt.Value, product.Quantity,
		product.Active,
	)
}
