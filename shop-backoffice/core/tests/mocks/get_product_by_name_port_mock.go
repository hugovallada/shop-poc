package mocks

import (
	"context"
	"errors"
	"strings"

	"github.com/hugovallada/shop-poc/shop-backoffice/core/ports/out/dto"
)

type GetProductByNameOutputPortMock struct{}

func (g *GetProductByNameOutputPortMock) Execute(p0 context.Context, p1 string) ([]dto.GetProductByNameResponse, error) {
	if strings.HasPrefix(p1, "ERROR") {
		return nil, errors.New("cant search with gsi")
	} else if strings.HasPrefix(p1, "MULTI") {
		return []dto.GetProductByNameResponse{&GetProductResponseMock{}, &GetProductResponseMock{}}, nil
	} else if strings.HasPrefix(p1, "Product") {
		return []dto.GetProductByNameResponse{&GetProductResponseMock{}}, nil
	} else {
		return []dto.GetProductByNameResponse{}, nil
	}
}
