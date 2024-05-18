package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/shop-poc/shop-backoffice/core"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/tests/mocks"
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
	/*productRepository := data.NewProductRepository(*config.ConfigDynamoDB())
	persistAdapter := out.NewPersistProductDynamoOutputAdapter(productRepository) */ //TODO: Usar esse bloco para substituir o mock
	persistPort := mocks.PersistProductOutputPortMock{}                              //TODO: Replace Mock
	createProductUseCase := core.NewCreateProductUseCase(&persistPort)
	return controller.NewCreateProductController(createProductUseCase)
}

func initHealthCheckController() controller.HealthCheckController {
	return controller.HealthCheckController{}
}

func initDependencies() controller.CreateProductController {
	createProductController := initCreateProductController()
	return createProductController
}
