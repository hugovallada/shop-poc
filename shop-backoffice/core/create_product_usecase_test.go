package core

import (
	"context"
	"testing"

	"github.com/hugovallada/shop-poc/shop-backoffice/core/customerror"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/tests/mocks"
	"github.com/stretchr/testify/assert"
)

var cp CreateProductUseCase = CreateProductUseCase{
	persistProductOutputPort:   &mocks.PersistProductOutputPortMock{},
	getProductByNameOutputPort: &mocks.GetProductByNameOutputPortMock{},
}

func TestCreateProductUseCaseJustRuns(t *testing.T) {
	error := cp.Execute(context.Background(), &mocks.CreateProductParameterMock{})
	assert.Nil(t, error)
}

func TestCreateProductUseCaseJustRunsWhenExactlyOneProductIsFound(t *testing.T) {
	error := cp.Execute(context.Background(), &mocks.CreateProductParameterMock{ProductName: "Product"})
	assert.Nil(t, error)
}

func TestCreateProductUseCaseWithError(t *testing.T) {
	error := cp.Execute(context.Background(), &mocks.CreateProductParameterMock{ProductName: "ERROR"})
	assert.NotNil(t, error)
}

func TestCreateProductUseCaseShouldReturnAnErrorWhenMoreThenOneProductIsFound(t *testing.T) {
	error := cp.Execute(context.Background(), &mocks.CreateProductParameterMock{ProductName: "MULTIPLE"})
	assert.NotNil(t, error)
	assert.ErrorIs(t, error, customerror.NewInternalError("quantidade de dados pré existentes é inválida"))
}
