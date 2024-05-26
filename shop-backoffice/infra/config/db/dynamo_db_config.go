package db

import (
	"log/slog"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ConfigDynamoDB() *dynamodb.DynamoDB {
	return dynamodb.New(prepareSess())
}

func ConfigDynamoDBLocal() *dynamodb.DynamoDB {
	dynamoClient := dynamodb.New(prepareLocalSess())
	_, err := dynamoClient.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String("Products"),
	})
	if err != nil {
		_, err := dynamoClient.CreateTable(&dynamodb.CreateTableInput{
			TableName: aws.String("Products"),
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("id"),
					KeyType:       aws.String("HASH"),
				},
			},
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("id"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("name"),
					AttributeType: aws.String("S"),
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(5),
				WriteCapacityUnits: aws.Int64(5),
			},
			GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
				{
					IndexName: aws.String("name_gsi"),
					KeySchema: []*dynamodb.KeySchemaElement{
						{
							AttributeName: aws.String("name"),
							KeyType:       aws.String("HASH"),
						},
					},
					Projection: &dynamodb.Projection{
						ProjectionType: aws.String("ALL"),
					},
					ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
						ReadCapacityUnits:  aws.Int64(5),
						WriteCapacityUnits: aws.Int64(5),
					},
				},
			},
		})
		if err != nil {
			slog.Error("Não foi possível configurar a tabela, será necessário criar manualmente", slog.Any("error", err.Error()))
		}
	}
	return dynamoClient
}

func prepareLocalSess() *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("sa-east-1"),
		Endpoint: aws.String("http://localhost:8000"),
	}))
}

func prepareSess() *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Region: aws.String("sa-east-1"),
	}))
}
