package dto

type CreateProductRequest struct {
	Name       string   `json:"nome" binding:"required,min=4,max=20"`
	Department string   `json:"departamento" binding:"required,min=5,max=30"`
	Tags       []string `json:"tags" binding:"required"`
	Price      uint64   `json:"preco" binding:"required,gt=0"`
	Quantity   uint8    `json:"quantidade" binding:"required,lte=100"`
	Activate   bool     `json:"ativo"`
}

func (cp CreateProductRequest) Validate() []string {
	var errors []string
	if cp.Name == "" {
		errors = append(errors, "Name cant be empty")
	} else if len(cp.Name) < 2 {
		errors = append(errors, "Name needs to have more than 2 letters")
	}
	return errors
}

func (c CreateProductRequest) GetName() string {
	return c.Name
}

func (c CreateProductRequest) GetDepartment() string {
	return c.Department
}

func (c CreateProductRequest) GetTags() []string {
	return c.Tags
}

func (c CreateProductRequest) GetPrice() uint64 {
	return c.Price
}

func (c CreateProductRequest) GetQuantity() uint8 {
	return c.Quantity
}

func (c CreateProductRequest) ShouldActivate() bool {
	return c.Activate
}
