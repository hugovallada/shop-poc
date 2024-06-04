package mocks

type CreateProductParameterMock struct {
	ProductName string
	Price       uint64
}

func (c *CreateProductParameterMock) GetName() string {
	return c.ProductName
}

func (c *CreateProductParameterMock) GetDepartment() string {
	return "Department"
}

func (c *CreateProductParameterMock) GetTags() []string {
	return []string{"Tags, Full"}
}

func (c *CreateProductParameterMock) GetPrice() uint64 {
	if c.Price == 0 {
		return 1900
	}
	return c.Price
}

func (c *CreateProductParameterMock) GetQuantity() uint8 {
	return 5
}

func (c *CreateProductParameterMock) ShouldActivate() bool {
	return true
}
