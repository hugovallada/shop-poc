package model

import "github.com/hugovallada/shop-poc/shop-backoffice/internal/core/ports/out/dto"

type ProductModel struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Department string   `json:"department"`
	Tags       []string `json:"tags"`
	Price      uint64   `json:"price"`
	Quantity   uint8    `json:"quantity"`
	TotalPrice uint64   `json:"total_price"`
	Active     bool     `json:"active"`
	CreatedAt  uint64   `json:"created_at"`
	UpdatedAt  uint64   `json:"updated_at"`
}

func NewProductModel(interfaceProductParamter dto.PersistProductParameter) ProductModel {

	return ProductModel{
		ID:         interfaceProductParamter.GetID().String(),
		Name:       interfaceProductParamter.GetName(),
		Department: interfaceProductParamter.GetDepartment(),
		Tags:       interfaceProductParamter.GetTags(),
		Price:      interfaceProductParamter.GetPrice(),
		Quantity:   interfaceProductParamter.GetQuantity(),
		TotalPrice: interfaceProductParamter.GetTotalPrice(),
		Active:     interfaceProductParamter.IsActive(),
		CreatedAt:  interfaceProductParamter.GetCreatedAt(),
		UpdatedAt:  interfaceProductParamter.GetUpdatedAt(),
	}
}
