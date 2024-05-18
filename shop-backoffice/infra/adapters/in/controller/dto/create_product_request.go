package dto

type CreateProductRequest struct {
	Name           string   `json:"nome"`
	Department     string   `json:"departamento"`
	Tags           []string `json:"tags"`
	Price          uint64   `json:"preco"`
	Quantity       uint8    `json:"quantidade"`
	ShouldActivate bool     `json:"ativo"`
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
