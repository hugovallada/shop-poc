package dto

import "github.com/google/uuid"

type GetProductByNameResponse interface {
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
