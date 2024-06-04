package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/adapters/in/controller"
)

func InitRoutes(r *gin.RouterGroup, cp controller.CreateProductController) {
	r.POST("/", cp.CreateProduct)
}

func InitActuatorRoutes(r *gin.RouterGroup, hc controller.HealthCheckController) {
	r.GET("/status", hc.HealthStatus)
}
