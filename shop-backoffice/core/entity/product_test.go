package entity

import (
	"testing"

	"github.com/hugovallada/shop-poc/shop-backoffice/core/vo"
	"github.com/stretchr/testify/assert"
)

func TestEnableProductOffering(t *testing.T) {
	product := Product{
		Active: false,
	}
	product.EnableProductOffering()
	assert.True(t, product.Active)
}

func TestDisableProductOffering(t *testing.T) {
	product := Product{
		Active: true,
	}
	product.DisableProductOffering()
	assert.False(t, product.Active)
}

func TestGivenAProductWith0ProductsAndStatusActiveAndWithAnTotalPriceBiggerThanZeroWhenTheMethodUpdateProductISCalledThenDisableTheProductAndSetTotalPriceToZero(t *testing.T) {
	product := Product{
		Active:     true,
		Quantity:   0,
		TotalPrice: vo.NewMoney(600000),
	}
	product.UpdateProduct()
	assert.False(t, product.Active)
	assert.EqualValues(t, 0, product.TotalPrice.Value)
}

func TestGivenAProductWith5ProductsAndStatusActiveAndWithAnTotalPriceBiggerThanZeroWhenTheMethodUpdateProductISCalledThenSetTotalPriceToTheRightValue(t *testing.T) {
	product := Product{
		Active:     true,
		Quantity:   5,
		Price:      vo.NewMoney(20000),
		TotalPrice: vo.NewMoney(20000),
	}
	product.UpdateProduct()
	assert.True(t, product.Active)
	assert.EqualValues(t, 100000, product.TotalPrice.Value)
}
