package validations

import (
	"github.com/hugovallada/shop-poc/shop-backoffice/core/customerror"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/entity"
)

type ProductValidation struct {
}

func (pv ProductValidation) Validate(product entity.Product) customerror.ValidationError {
	return customerror.ValidationError{}
}
