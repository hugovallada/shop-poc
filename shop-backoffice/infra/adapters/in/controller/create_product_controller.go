package controller

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct{}

func (cp CreateProductController) CreateProduct(c *gin.Context) {
	var mapa map[string]any
	err := c.BindJSON(&mapa)
	if err != nil {
		c.JSON(500, "Erro interno ao fazer a conversão de variaveis")
		return
	}
	slog.Info("Cadastrando produto", slog.Any("produto", mapa))
	c.JSON(201, "Created")
}
