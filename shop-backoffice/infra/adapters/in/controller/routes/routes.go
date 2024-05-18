package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/in/controller"
)

func InitRoutes(r *gin.RouterGroup, cp controller.CreateProductController) {
	r.POST("/products", cp.CreateProduct)
}
