package controller

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/correlationcontexthandler"

	"github.com/hugovallada/shop-poc/shop-backoffice/internal/core/customerror"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/core/ports/in"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/adapters/in/controller/dto"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/utils"
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
	var productRequest dto.CreateProductRequest
	ctx, err := contextWithHeadersValues(context.Background(), c)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.BindJSON(&productRequest)
	if err != nil {
		slog.ErrorContext(ctx, "error while deserializing", slog.Any("error", err.Error()))
		validationErrors := utils.ValidateErrors(err, dto.CreateProductRequest{})
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}
	slog.InfoContext(ctx, "Cadastrando produto", slog.Any("produto", productRequest))
	if err = cp.createProductUseCase.Execute(ctx, productRequest); err != nil {
		httpStatus := getErrorHttpStatus(err)
		c.JSON(httpStatus, gin.H{"error": err.Error()})
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
	flowId := getHeaderValue("flowId", c)
	ctx := context.WithValue(parentContext, correlationcontexthandler.CORRELATION_ID, correlationId)
	ctx = context.WithValue(ctx, correlationcontexthandler.TRACE_ID, traceId)
	ctx = context.WithValue(ctx, correlationcontexthandler.FLOW_ID, flowId)
	return ctx, nil
}

func getHeaderValue(key string, c *gin.Context) string {
	return strings.Trim(strings.Trim(c.GetHeader(key), "\""), "")
}

func getErrorHttpStatus(err error) int {
	switch err.(type) {
	case customerror.InternalError:
		return http.StatusInternalServerError
	case customerror.UnprocessableEntityError:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusBadGateway
	}
}
