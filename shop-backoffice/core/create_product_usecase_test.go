package core

import (
	"testing"

	"github.com/hugovallada/shop-poc/shop-backoffice/core/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductUseCaseJustRuns(t *testing.T) {
	cp := CreateProductUseCase{
		persistProductOutputPort: &mocks.PersistProductOutputPortMock{},
	}
	error := cp.Execute(&mocks.CreateProductParameterMock{})
	assert.Nil(t, error)
}

func TestCreateProductUseCaseWithError(t *testing.T) {
	cp := CreateProductUseCase{
		persistProductOutputPort: &mocks.PersistProductOutputPortMock{},
	}
	error := cp.Execute(&mocks.CreateProductParameterMock{IsError: true})
	assert.NotNil(t, error)
}
