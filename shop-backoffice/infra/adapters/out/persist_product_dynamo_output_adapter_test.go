package out

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/hugovallada/shop-poc/shop-backoffice/core/ports/out/dto"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewPersistProductDynamoOutputAdapterSuccess(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repository := mocks.NewMockProductRepositoryInterface(control)
	sut := NewPersistProductDynamoOutputAdapter(repository)

	repository.EXPECT().SaveProduct(gomock.Any(), gomock.Any()).Return(nil)

	err := sut.Execute(context.Background(), dto.NewCreateProductParameter(
		uuid.New(), "produto", "departamento", []string{},
		82999, 829990, uint64(time.Now().Unix()), uint64(time.Now().Unix()),
		10, true,
	))

	assert.Nil(t, err)
}

func TestNewPersistProductDynamoOutputAdapterReturnsAnErrorWhenRepositoryCallFails(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repository := mocks.NewMockProductRepositoryInterface(control)
	sut := NewPersistProductDynamoOutputAdapter(repository)

	repository.EXPECT().SaveProduct(gomock.Any(), gomock.Any()).Return(errors.New("invalid payload"))

	err := sut.Execute(context.Background(), dto.NewCreateProductParameter(
		uuid.New(), "produto", "departamento", []string{},
		82999, 829990, uint64(time.Now().Unix()), uint64(time.Now().Unix()),
		10, true,
	))

	assert.NotNil(t, err)
	assert.Equal(t, "invalid payload", err.Error())
}
