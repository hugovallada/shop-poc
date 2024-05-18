package data

import (
	"log/slog"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/hugovallada/shop-poc/shop-backoffice/infra/data/model"
)

const (
	TABLE_NAME = "tbx_000"
)

type ProductRepository struct {
	dynamo dynamodb.DynamoDB
}

func NewProductRepository(dynamo dynamodb.DynamoDB) ProductRepository {
	return ProductRepository{
		dynamo: dynamo,
	}
}

func (pr ProductRepository) SaveProduct(productModel model.ProductModel) error {
	dynamoItem, err := dynamodbattribute.MarshalMap(productModel)
	if err != nil {
		slog.Error("Can't desserialize product model into dynamo map", slog.Any("model", productModel))
		return err
	}
	_, err = pr.dynamo.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item:      dynamoItem,
	})
	return err
}
