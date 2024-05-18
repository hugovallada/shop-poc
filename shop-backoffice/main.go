package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/in/controller"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/in/controller/routes"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

}

func main() {
	slog.Info("Inicializando a aplicação")
	router := gin.Default()
	controller := controller.CreateProductController{}
	contextGroup := router.Group("/backoffice")
	productsGroup := contextGroup.Group("/products")
	routes.InitRoutes(productsGroup, controller)
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
