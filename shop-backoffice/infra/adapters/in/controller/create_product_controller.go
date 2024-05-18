package controller

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	inputAdapterDto "github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/in/controller/dto"
)

type CreateProductController struct{}

func (cp CreateProductController) CreateProduct(c *gin.Context) {
	var productRequest inputAdapterDto.CreateProductRequest
	err := c.BindJSON(&productRequest)
	if err != nil {
		c.JSON(500, "Erro interno ao fazer a conversão de variaveis")
		return
	}
	fmt.Println(productRequest)
	slog.Info("Cadastrando produto", slog.Any("produto", productRequest))
	c.JSON(201, "Created")
}
