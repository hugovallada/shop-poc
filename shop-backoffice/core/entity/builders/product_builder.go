package builders

import (
	"github.com/hugovallada/shop-poc/shop-backoffice/core/entity"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/vo"
)

type ProductBuilder struct {
	Name       string
	Department string
	Tags       []string
	Price      uint64
	Quantity   uint8
	Active     bool
	UpdatedAt  uint64
}

func NewProductBuilder() *ProductBuilder {
	return &ProductBuilder{}
}

func (pb *ProductBuilder) SetName(name string) *ProductBuilder {
	pb.Name = name
	return pb
}

func (pb *ProductBuilder) SetDepartment(department string) *ProductBuilder {
	pb.Department = department
	return pb
}

func (pb *ProductBuilder) SetTags(tags []string) *ProductBuilder {
	pb.Tags = tags
	return pb
}

func (pb *ProductBuilder) SetPrice(price uint64) *ProductBuilder {
	pb.Price = price
	return pb
}

func (pb *ProductBuilder) SetQuantity(quantity uint8) *ProductBuilder {
	pb.Quantity = quantity
	return pb
}

func (pb *ProductBuilder) SetActive(active bool) *ProductBuilder {
	pb.Active = active
	return pb
}

func (pb *ProductBuilder) SetUpdatedAt(updatedAt uint64) *ProductBuilder {
	pb.UpdatedAt = updatedAt
	return pb
}

func (pb *ProductBuilder) Build() entity.Product {
	return entity.Product{
		ID:         vo.NewID(),
		Name:       vo.NewName(pb.Name),
		Department: vo.NewDepartment(pb.Department),
		Tags:       vo.NewTags(pb.Tags),
		Price:      vo.NewMoney(pb.Price),
		Quantity:   pb.Quantity,
		TotalPrice: vo.NewMoney(pb.Price * uint64(pb.Quantity)),
		Active:     pb.Active,
		CreatedAt:  vo.NewTimestamp(),
		UpdatedAt:  vo.NewTimestamp(),
	}
}
