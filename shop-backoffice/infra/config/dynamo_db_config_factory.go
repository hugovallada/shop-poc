package config

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func BuildDynamoDBConfig(environment string) *dynamodb.DynamoDB {
	if environment == "local" {
		return ConfigDynamoDBLocal()
	} else {
		return ConfigDynamoDB()
	}
}
