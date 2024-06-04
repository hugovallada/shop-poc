package core

import (
	"context"
	"log/slog"

	"github.com/hugovallada/shop-poc/shop-backoffice/internal/core/customerror"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/core/entity/factorys"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/core/ports/in/dto"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/core/ports/out"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/core/translators"
)

type CreateProductUseCase struct {
	persistProductOutputPort   out.PersistProductOutputPort
	getProductByNameOutputPort out.GetProductByNameOutputPort
}

func NewCreateProductUseCase(
	persistProductPort out.PersistProductOutputPort,
	getProductByNameOutputPort out.GetProductByNameOutputPort,
) CreateProductUseCase {
	return CreateProductUseCase{
		persistProductOutputPort:   persistProductPort,
		getProductByNameOutputPort: getProductByNameOutputPort,
	}
}

func (cp CreateProductUseCase) Execute(ctx context.Context, createProductParameter dto.CreateProductParameter) error {
	product := factorys.ProductFactoryFromCreateProductParameter(createProductParameter)
	product.UpdateProduct()
	slog.InfoContext(ctx, "Product updated")
	products, err := cp.getProductByNameOutputPort.Execute(ctx, product.Name.Value)
	if err != nil {
		slog.ErrorContext(ctx, "erro na execução do usecase, na chamada de busca de produtos", slog.Any("erro", err.Error()))
		return err
	}
	if len(products) > 1 {
		slog.ErrorContext(ctx, "mais de um produto com o mesmo nome foi cadastrado na tabela, contate o administrador do sistema")
		return customerror.NewInternalError("quantidade de dados pré existentes é inválida")
	}
	if len(products) == 1 {
		productEntity := factorys.ProductFactoryFromGetProuctByNameResponse(products[0])
		if updated := product.UpdateProductWithExistingData(productEntity); !updated {
			slog.ErrorContext(ctx, "Failed to update product")
			return customerror.NewUnprocessableEntityError("there is no new info to update")
		}
		slog.InfoContext(ctx, "product updated with new data", slog.Any("oldProduct", productEntity), slog.Any("newProduct", product))
	}
	product.UpdateProduct()
	persistProductParameter := translators.FromProductEntityToPersistProductParameter(product)
	return cp.persistProductOutputPort.Execute(ctx, persistProductParameter)
}
