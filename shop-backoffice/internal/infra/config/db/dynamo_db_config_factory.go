package db

import (
	"os"
	"slices"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func BuildDynamoDBConfig() *dynamodb.DynamoDB {
	cloudEnvironments := []string{"dev", "hom", "prod"}
	environment := os.Getenv("ENVIRONMENT")
	if !slices.Contains(cloudEnvironments, environment) {
		return ConfigDynamoDBLocal()
	} else {
		return ConfigDynamoDB()
	}
}
