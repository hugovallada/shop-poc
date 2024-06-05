package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/core"
	outputPort "github.com/hugovallada/shop-poc/shop-backoffice/internal/core/ports/out"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/core/tests/mocks"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/adapters/in/controller"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/adapters/in/controller/middlewares"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/adapters/in/controller/routes"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/adapters/out"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/config/db"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/config/logs"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/data"
	"github.com/spf13/viper"
)

var (
	env string
)

func init() {
	logs.SetDefaultSlogHandler()
	initViperConfig()
	env = os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "local"
	}
}

func initViperConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("properties")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Can't load config")
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
