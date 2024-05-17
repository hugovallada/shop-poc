package mocks

type CreateProductParameterMock struct {
	IsError bool
}

func (c *CreateProductParameterMock) GetName() string {
	if c.IsError {
		return "FAILURE_PRODUCT"
	}
	return "Product"
}

func (c *CreateProductParameterMock) GetDepartment() string {
	return "Department"
}

func (c *CreateProductParameterMock) GetTags() []string {
	return []string{"Tags, Full"}
}

func (c *CreateProductParameterMock) GetPrice() uint64 {
	return 1900
}

func (c *CreateProductParameterMock) GetQuantity() uint8 {
	return 5
}

func (c *CreateProductParameterMock) ShouldActivate() bool {
	return true
}
