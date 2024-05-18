package main

import (
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/in/controller"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/in/controller/routes"
)

func main() {
	slog.Info("Inicializando a aplicação")
	router := gin.Default()
	controller := controller.CreateProductController{}
	routes.InitRoutes(&router.RouterGroup, controller)
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
