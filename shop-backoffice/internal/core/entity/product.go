package entity

import (
	"slices"

	"github.com/hugovallada/shop-poc/shop-backoffice/internal/core/vo"
)

type Product struct {
	ID         vo.ID
	Name       vo.Name
	Department vo.Department
	Tags       vo.Tags
	Price      vo.Money
	Quantity   uint8
	TotalPrice vo.Money
	Active     bool
	CreatedAt  vo.Timestamp
	UpdatedAt  vo.Timestamp
}

func (p *Product) UpdateProductWithNewData(newProductData Product) {
	p.ID = newProductData.ID
	p.Department = newProductData.Department
	p.Tags = newProductData.Tags
	p.Price = newProductData.Price
	p.Quantity = newProductData.Quantity
	p.Active = newProductData.Active
}

func (p *Product) UpdateProductWithExistingData(oldProduct Product) bool {
	if p.Quantity == oldProduct.Quantity && p.Price == oldProduct.Price && p.Active == oldProduct.Active && p.Department == oldProduct.Department && len(p.Tags.Value) == len(oldProduct.Tags.Value) {
		var sameTags = true
		for _, tag := range p.Tags.Value {
			if !slices.Contains(oldProduct.Tags.Value, tag) {
				sameTags = false
			}
		}
		if sameTags {
			return false
		}
	}
	p.ID = oldProduct.ID
	p.CreatedAt = oldProduct.CreatedAt
	return true
}

func (p *Product) UpdateProduct() {
	p.UpdatedAt = vo.NewTimestamp()
	if p.Quantity == 0 {
		p.DisableProductOffering()
		p.TotalPrice = vo.NewMoney(0)
	} else {
		p.UpdateTotalPrice()
	}
}

func (p *Product) UpdateTotalPrice() {
	p.TotalPrice = vo.NewMoney(p.Price.Value * uint64(p.Quantity))
}

func (p *Product) DisableProductOffering() {
	p.Active = false
}

func (p *Product) EnableProductOffering() {
	p.Active = true
}
