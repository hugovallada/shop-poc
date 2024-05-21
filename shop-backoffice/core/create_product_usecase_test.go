package core

import (
	"context"
	"testing"

	"github.com/hugovallada/shop-poc/shop-backoffice/core/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductUseCaseJustRuns(t *testing.T) {
	cp := CreateProductUseCase{
		persistProductOutputPort: &mocks.PersistProductOutputPortMock{},
	}
	error := cp.Execute(context.Background(), &mocks.CreateProductParameterMock{})
	assert.Nil(t, error)
}

func TestCreateProductUseCaseWithError(t *testing.T) {
	cp := CreateProductUseCase{
		persistProductOutputPort: &mocks.PersistProductOutputPortMock{},
	}
	error := cp.Execute(context.Background(), &mocks.CreateProductParameterMock{IsError: true})
	assert.NotNil(t, error)
}
