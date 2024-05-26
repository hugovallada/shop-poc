package dto

import "github.com/google/uuid"

type GetProductByNameResponseImpl struct {
	ID         uuid.UUID
	Name       string
	Department string
	Tags       []string
	Price      uint64
	Quantity   uint8
	TotalPrice uint64
	Active     bool
	CreatedAt  uint64
	UpdatedAt  uint64
}

func (g GetProductByNameResponseImpl) GetName() string {
	return g.Name
}

func (g GetProductByNameResponseImpl) GetDepartment() string {
	return g.Department
}

func (g GetProductByNameResponseImpl) GetTags() []string {
	return g.Tags
}

func (g GetProductByNameResponseImpl) GetPrice() uint64 {
	return g.Price
}

func (g GetProductByNameResponseImpl) GetQuantity() uint8 {
	return g.Quantity
}

func (g GetProductByNameResponseImpl) GetID() uuid.UUID {
	return g.ID
}

func (g GetProductByNameResponseImpl) GetTotalPrice() uint64 {
	return g.TotalPrice
}

func (g GetProductByNameResponseImpl) IsActive() bool {
	return g.Active
}

func (g GetProductByNameResponseImpl) GetCreatedAt() uint64 {
	return g.CreatedAt
}

func (g GetProductByNameResponseImpl) GetUpdatedAt() uint64 {
	return g.UpdatedAt
}
