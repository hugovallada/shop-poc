package dto

import "github.com/google/uuid"

type PersistProductParameter interface {
	GetID() uuid.UUID
	GetName() string
	GetDepartment() string
	GetTags() []string
	GetPrice() uint64
	GetQuantity() uint8
	GetTotalPrice() uint64
	IsActive() bool
	GetCreatedAt() uint64
	GetUpdatedAt() uint64
}

type persistProductParameterImpl struct {
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

func NewCreateProductParameter(
	id uuid.UUID,
	name, department string,
	tags []string,
	price, totalPrice, createdAt, updatedAt uint64,
	quantity uint8,
	isActive bool,
) persistProductParameterImpl {
	return persistProductParameterImpl{
		ID: id, Name: name, Department: department, Tags: tags,
		Price: price, Quantity: quantity, TotalPrice: totalPrice,
		Active: isActive, CreatedAt: createdAt, UpdatedAt: updatedAt,
	}
}

func (p persistProductParameterImpl) GetID() uuid.UUID {
	return p.ID
}

func (p persistProductParameterImpl) GetName() string {
	return p.Name
}

func (p persistProductParameterImpl) GetDepartment() string {
	return p.Department
}

func (p persistProductParameterImpl) GetTags() []string {
	return p.Tags
}

func (p persistProductParameterImpl) GetPrice() uint64 {
	return p.Price
}

func (p persistProductParameterImpl) GetQuantity() uint8 {
	return p.Quantity
}

func (p persistProductParameterImpl) GetTotalPrice() uint64 {
	return p.TotalPrice
}

func (p persistProductParameterImpl) IsActive() bool {
	return p.Active
}

func (p persistProductParameterImpl) GetCreatedAt() uint64 {
	return p.CreatedAt
}

func (p persistProductParameterImpl) GetUpdatedAt() uint64 {
	return p.UpdatedAt
}
