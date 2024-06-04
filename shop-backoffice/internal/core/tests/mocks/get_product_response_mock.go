package mocks

import (
	"time"

	"github.com/google/uuid"
)

type GetProductResponseMock struct {
}

func (g *GetProductResponseMock) GetName() string {
	return "Product"
}

func (g *GetProductResponseMock) GetDepartment() string {
	return "Department"
}

func (g *GetProductResponseMock) GetTags() []string {
	return []string{"Tags, Full"}
}

func (g *GetProductResponseMock) GetPrice() uint64 {
	return 1900
}

func (g *GetProductResponseMock) GetQuantity() uint8 {
	return 5
}

func (g *GetProductResponseMock) ShouldActivate() bool {
	return true
}

func (g *GetProductResponseMock) GetID() uuid.UUID {
	return uuid.New()
}

func (g *GetProductResponseMock) GetTotalPrice() uint64 {
	return uint64(g.GetQuantity()) * g.GetPrice()
}

func (g *GetProductResponseMock) IsActive() bool {
	return true
}

func (g *GetProductResponseMock) GetCreatedAt() uint64 {
	return uint64(time.Now().Unix())
}

func (g *GetProductResponseMock) GetUpdatedAt() uint64 {
	return uint64(time.Now().Unix())
}
