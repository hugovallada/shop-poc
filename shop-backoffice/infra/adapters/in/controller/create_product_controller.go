package controller

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct{}

func (cp CreateProductController) CreateProduct(c *gin.Context) {
	var mapa map[string]any
	c.BindJSON(&mapa)
	slog.Info("Cadastrando produto", slog.Any("produto", mapa))
	c.JSON(201, "Created")
}
