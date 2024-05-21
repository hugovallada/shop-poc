package dto

type CreateProductParameter interface {
	GetName() string
	GetDepartment() string
	GetTags() []string
	GetPrice() uint64
	GetQuantity() uint8
	ShouldActivate() bool
}
