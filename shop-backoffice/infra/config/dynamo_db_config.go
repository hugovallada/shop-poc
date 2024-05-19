package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ConfigDynamoDB() *dynamodb.DynamoDB {
	return dynamodb.New(prepareSess())
}

func ConfigDynamoDBLocal() *dynamodb.DynamoDB {
	return dynamodb.New(prepareLocalSess())
}

func prepareLocalSess() *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("sa-east-1"),
		Endpoint: aws.String("http://localhost:8000"),
	}))
}

func prepareSess() *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("sa-east-1"),
		Endpoint: aws.String("http://localhost:8000"),
	}))
}
