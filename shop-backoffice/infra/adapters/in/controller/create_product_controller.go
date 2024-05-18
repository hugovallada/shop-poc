package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct{}

func (cp CreateProductController) CreateProduct(c *gin.Context) {
	var mapa map[string]any
	c.BindJSON(&mapa)
	fmt.Println(mapa)
	c.JSON(201, "Created")
}
