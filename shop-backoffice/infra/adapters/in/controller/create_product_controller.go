package controller

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/ports/in"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/utils/shared"
	inputAdapterDto "github.com/hugovallada/shop-poc/shop-backoffice/infra/adapters/in/controller/dto"
)

type CreateProductController struct {
	createProductUseCase in.CreateProductUseCaseInputPort
}

func NewCreateProductController(createProductUseCase in.CreateProductUseCaseInputPort) CreateProductController {
	return CreateProductController{
		createProductUseCase: createProductUseCase,
	}
}

func (cp CreateProductController) CreateProduct(c *gin.Context) {
	var productRequest inputAdapterDto.CreateProductRequest
	ctx, err := contextWithHeadersValues(context.Background(), c)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.BindJSON(&productRequest)
	if err != nil {
		slog.ErrorContext(ctx, "error while deserializing")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while deserializing"})
		return
	}
	if errors := productRequest.Validate(); len(errors) > 0 {
		slog.ErrorContext(ctx, "validation error", slog.Any("errors", errors))
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}
	slog.InfoContext(ctx, "Cadastrando produto", slog.Any("produto", productRequest))
	if err = cp.createProductUseCase.Execute(ctx, productRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't persist the product"})
		slog.ErrorContext(ctx, "Can't persist the product", slog.Any("product", productRequest))
		return
	}
	c.JSON(201, "Created")
}

func contextWithHeadersValues(parentContext context.Context, c *gin.Context) (context.Context, error) {
	correlationId := getHeaderValue("correlationId", c)
	if correlationId == "" {
		return nil, errors.New("correlation id can't be null")
	}
	traceId := getHeaderValue("traceId", c)
	ctx := context.WithValue(parentContext, shared.CORRELATION_ID, correlationId)
	ctx = context.WithValue(ctx, shared.TRACE_ID, traceId)
	return ctx, nil
}

func getHeaderValue(key string, c *gin.Context) string {
	return strings.Trim(strings.Trim(c.GetHeader(key), "\""), "")
}
