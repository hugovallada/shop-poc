package data

import (
	"context"
	"log/slog"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/hugovallada/shop-poc/shop-backoffice/internal/infra/data/model"
	"github.com/spf13/viper"
)

type ProductRepository struct {
	dynamo    dynamodb.DynamoDB
	tableName string
	gsiName   string
}

func NewProductRepository(dynamo dynamodb.DynamoDB) ProductRepository {
	dynamoConfig := viper.GetStringMapString("dynamo")
	return ProductRepository{
		dynamo:    dynamo,
		tableName: dynamoConfig["table_name"],
		gsiName:   dynamoConfig["gsi_name"],
	}
}

func (pr ProductRepository) SaveProduct(ctx context.Context, productModel model.ProductModel) error {
	dynamoItem, err := dynamodbattribute.MarshalMap(productModel)
	if err != nil {
		slog.Error("Can't desserialize product model into dynamo map", slog.Any("model", productModel))
		return err
	}
	slog.InfoContext(ctx, "successfully marshaled productModel into dynamoItemMap", slog.Any("dynamoItem", dynamoItem))
	_, err = pr.dynamo.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(pr.tableName),
		Item:      dynamoItem,
	})
	if err != nil {
		slog.ErrorContext(ctx, "Error while creating product", slog.Any("error", err))
	}
	return err
}

func (pr ProductRepository) GetProductsByName(ctx context.Context, name string) ([]model.ProductModel, error) {
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(pr.tableName),
		IndexName: aws.String(pr.gsiName),
		KeyConditions: map[string]*dynamodb.Condition{
			"name": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(name),
					},
				},
			},
		},
		Limit: aws.Int64(2),
	}
	queryResult, err := pr.dynamo.Query(queryInput)
	if err != nil {
		slog.ErrorContext(ctx, "Can't query by secondary index", slog.Any("queryInput", queryInput), slog.Any("error", err))
		return nil, err
	}
	var products []model.ProductModel
	if err = dynamodbattribute.UnmarshalListOfMaps(queryResult.Items, &products); err != nil {
		slog.ErrorContext(ctx, "can't desserialize items", slog.Any("items", queryResult.Items))
		return nil, err
	}
	slog.InfoContext(ctx, "Sucessfully returned items", slog.Int("items returned", len(products)), slog.Any("products", products))
	return products, nil
}
