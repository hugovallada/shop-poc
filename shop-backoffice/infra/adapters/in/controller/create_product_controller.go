package controller

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	inputAdapterDto "github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/in/controller/dto"
)

type CreateProductController struct{}

func (cp CreateProductController) CreateProduct(c *gin.Context) {
	var productRequest inputAdapterDto.CreateProductRequest
	err := c.BindJSON(&productRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while deserializing"})
		return
	}
	if errors := productRequest.Validate(); len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}
	slog.Info("Cadastrando produto", slog.Any("produto", productRequest))
	c.JSON(201, "Created")
}
