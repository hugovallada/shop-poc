package factorys

import (
	"github.com/hugovallada/shop-poc/shop-backoffice/core/entity"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/entity/builders"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/ports/in/dto"
	outputDto "github.com/hugovallada/shop-poc/shop-backoffice/core/ports/out/dto"
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

func ProductFactoryFromGetProuctByNameResponse(getProductByNameResponse outputDto.GetProductByNameResponse) entity.Product {
	return builders.
		NewProductBuilder().
		SetID(getProductByNameResponse.GetID().String()).
		SetName(getProductByNameResponse.GetName()).
		SetDepartment(getProductByNameResponse.GetDepartment()).
		SetTags(getProductByNameResponse.GetTags()).
		SetPrice(getProductByNameResponse.GetPrice()).
		SetQuantity(getProductByNameResponse.GetQuantity()).
		SetActive(getProductByNameResponse.IsActive()).
		SetUpdatedAt(getProductByNameResponse.GetUpdatedAt()).
		SetCreatedAt(getProductByNameResponse.GetCreatedAt()).
		Build()
}
