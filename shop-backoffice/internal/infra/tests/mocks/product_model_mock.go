package mocks

import (
	"time"

	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/data/model"
)

func MockProductModel() model.ProductModel {
	return model.ProductModel{
		ID:         "33c82050-c62f-4ca3-9a91-dd6425dc1000",
		Name:       "Notebook",
		Department: "TI",
		Tags:       []string{"TI", "Eletronico"},
		Price:      uint64(100000),
		Quantity:   3,
		TotalPrice: uint64(300000),
		Active:     true,
		CreatedAt:  uint64(time.Now().Unix()),
		UpdatedAt:  uint64(time.Now().Unix()),
	}
}
