package factorys

import (
	"github.com/hugovallada/shop-poc/shop-backoffice/core/entity"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/entity/builders"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/ports/in/dto"
)

func ProductFactoryFromCreateProductParameter(createProductParameter dto.CreateProductParameter) entity.Product {
	return builders.
		NewProductBuilder().
		SetName(createProductParameter.GetName()).
		SetDepartment(createProductParameter.GetDepartment()).
		SetTags(createProductParameter.GetTags()).
		SetPrice(createProductParameter.GetPrice()).
		SetQuantity(createProductParameter.GetQuantity()).
		SetActive(createProductParameter.ShouldActivate()).
		Build()
}
