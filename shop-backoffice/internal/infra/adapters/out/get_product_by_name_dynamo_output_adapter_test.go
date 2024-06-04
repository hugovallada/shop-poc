package out

import (
	"context"
	"errors"
	"testing"

	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/data/model"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/tests/mocks"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/translators"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestShouldReturnAnEmptyListWhenNoItensAreFound(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	repository := mocks.NewMockProductRepositoryInterface(control)
	getAdapter := NewGetProductByNameOutputAdapter(repository)
	repository.EXPECT().GetProductsByName(gomock.Any(), gomock.Any()).Return([]model.ProductModel{}, nil)
	products, err := getAdapter.Execute(context.Background(), "Produto")
	assert.Nil(t, err)
	assert.Len(t, products, 0)
}

func TestShouldReturnAnListWithOneItem(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	repository := mocks.NewMockProductRepositoryInterface(control)
	getAdapter := NewGetProductByNameOutputAdapter(repository)
	mockedResponse := mocks.MockProductModel()
	expectedResponse := translators.FromProductModelToGetProductByNameResponseImpl(mockedResponse)
	repository.EXPECT().GetProductsByName(gomock.Any(), gomock.Any()).Return([]model.ProductModel{mockedResponse}, nil)
	products, err := getAdapter.Execute(context.Background(), "Produto")
	assert.Nil(t, err)
	assert.Len(t, products, 1)
	assert.NotEmpty(t, products)
	assert.Equal(t, expectedResponse, products[0])
}

func TestShouldReturnAnListWithMultipleItens(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	repository := mocks.NewMockProductRepositoryInterface(control)
	getAdapter := NewGetProductByNameOutputAdapter(repository)
	mockedResponse := mocks.MockProductModel()
	expectedResponse := translators.FromProductModelToGetProductByNameResponseImpl(mockedResponse)
	repository.EXPECT().GetProductsByName(gomock.Any(), gomock.Any()).Return([]model.ProductModel{mockedResponse, mockedResponse}, nil)
	products, err := getAdapter.Execute(context.Background(), "Produto")
	assert.Nil(t, err)
	assert.Len(t, products, 2)
	assert.NotEmpty(t, products)
	assert.Equal(t, expectedResponse, products[0])
}

func TestShouldReturnAnError(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	repository := mocks.NewMockProductRepositoryInterface(control)
	getAdapter := NewGetProductByNameOutputAdapter(repository)
	repository.EXPECT().GetProductsByName(gomock.Any(), gomock.Any()).Return(nil, errors.New("invalid gsi"))
	products, err := getAdapter.Execute(context.Background(), "Produto")
	assert.Nil(t, products)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid gsi")
}
