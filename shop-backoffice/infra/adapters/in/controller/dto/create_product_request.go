package dto

type CreateProductRequest struct {
	Name           string   `json:"nome"`
	Department     string   `json:"departamento"`
	Tags           []string `json:"tags"`
	Price          uint64   `json:"preco"`
	Quantity       uint8    `json:"quantidade"`
	ShouldActivate bool     `json:"ativo"`
}
