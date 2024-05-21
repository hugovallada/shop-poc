package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/shop-poc/shop-backoffice/core"
	outputPort "github.com/hugovallada/shop-poc/shop-backoffice/core/ports/out"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/tests/mocks"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/utils/shared"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/in/controller"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/in/controller/middlewares"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/in/controller/routes"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/out"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/config"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/data"
	"github.com/joho/godotenv"
)

func init() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	})
	keys := []any{shared.CORRELATION_ID, shared.TRACE_ID}
	ctxHandler := config.NewCorrelationMultiKey(keys, handler)
	logger := slog.New(ctxHandler)
	slog.SetDefault(logger)
	_ = godotenv.Load()
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
	env := os.Getenv("ENVIRONMENT")
	slog.Info("Initializing dependencies for environment " + env)
	var persistPort outputPort.PersistProductOutputPort
	productRepository := data.NewProductRepository(*config.BuildDynamoDBConfig(env))
	persistAdapter := out.NewPersistProductDynamoOutputAdapter(productRepository)
	persistMock := mocks.PersistProductOutputPortMock{}
	if env != "local" {
		persistPort = &persistAdapter
	} else {
		persistPort = &persistMock
	}
	createProductUseCase := core.NewCreateProductUseCase(persistPort)
	return controller.NewCreateProductController(createProductUseCase)
}

func initHealthCheckController() controller.HealthCheckController {
	return controller.HealthCheckController{}
}

func initDependencies() controller.CreateProductController {
	createProductController := initCreateProductController()
	return createProductController
}
