package db

import (
	"slices"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func BuildDynamoDBConfig(environment string) *dynamodb.DynamoDB {
	cloudEnvironments := []string{"dev", "hom", "prod"}
	if !slices.Contains(cloudEnvironments, environment) {
		return ConfigDynamoDBLocal()
	} else {
		return ConfigDynamoDB()
	}
}
