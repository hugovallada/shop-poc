package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/shop-poc/shop-backoffice/core"
	outputPort "github.com/hugovallada/shop-poc/shop-backoffice/core/ports/out"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/tests/mocks"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/in/controller"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/in/controller/middlewares"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/in/controller/routes"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/out"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/config/db"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/config/logs"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/data"
)

var (
	env string
)

func init() {
	logs.SetDefaultSlogHandler()
	env = os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "local"
	}
}

func main() {
	slog.Info("Inicializando a aplicação")
	router := gin.Default()
	router.Use(middlewares.CallDuration)

	contextGroup := router.Group("/backoffice")
	productsGroup := contextGroup.Group("/products")
	healthGroup := contextGroup.Group("/health")
	createProductController := initDependencies()
	routes.InitRoutes(productsGroup, createProductController)
	routes.InitActuatorRoutes(healthGroup, initHealthCheckController())
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}

func initCreateProductController() controller.CreateProductController {
	slog.Info("Initializing dependencies for environment " + env)
	var persistPort outputPort.PersistProductOutputPort
	productRepository := data.NewProductRepository(*db.BuildDynamoDBConfig(env))
	persistAdapter := out.NewPersistProductDynamoOutputAdapter(productRepository)
	getProductAdapter := out.NewGetProductByNameOutputAdapter(productRepository)
	persistMock := mocks.PersistProductOutputPortMock{}
	if env != "local" {
		persistPort = &persistAdapter
	} else {
		persistPort = &persistMock
	}
	createProductUseCase := core.NewCreateProductUseCase(persistPort, getProductAdapter)
	return controller.NewCreateProductController(createProductUseCase)
}

func initHealthCheckController() controller.HealthCheckController {
	return controller.HealthCheckController{}
}

func initDependencies() controller.CreateProductController {
	createProductController := initCreateProductController()
	return createProductController
}
