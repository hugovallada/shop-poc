package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/core"
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
	initViperConfig()
	logs.SetDefaultSlogHandler()
	env = os.Getenv("ENVIRONMENT")
	if env == "" {
		err := os.Setenv("ENVIRONMENT", "local")
		if err != nil {
			panic("Can't find a suitable environment")
		}
	}
}

func initViperConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("properties")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("can't load config", slog.String("error", err.Error()))
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
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			log.Fatal(err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	slog.Info("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v", err)
	}
	slog.Info("Server stopped")

}

func initCreateProductController() controller.CreateProductController {
	slog.Info("Initializing dependencies for environment " + env)
	productRepository := data.NewProductRepository(*db.BuildDynamoDBConfig())
	persistAdapter := out.NewPersistProductDynamoOutputAdapter(productRepository)
	getProductAdapter := out.NewGetProductByNameOutputAdapter(productRepository)
	createProductUseCase := core.NewCreateProductUseCase(&persistAdapter, getProductAdapter)
	return controller.NewCreateProductController(createProductUseCase)
}

func initHealthCheckController() controller.HealthCheckController {
	return controller.HealthCheckController{}
}

func initDependencies() controller.CreateProductController {
	createProductController := initCreateProductController()
	return createProductController
}
