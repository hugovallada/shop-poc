package entity

import (
	"github.com/hugovallada/shop-poc/shop-backoffice/core/vo"
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
	p.TotalPrice = vo.NewMoney(p.TotalPrice.Value * uint64(p.Quantity))
}

func (p *Product) DisableProductOffering() {
	p.Active = false
}

func (p *Product) EnableProductOffering() {
	p.Active = true
}
