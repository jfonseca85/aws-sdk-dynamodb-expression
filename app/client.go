package app

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jfonseca85/aws-sdk-dynamodb-expression/config"
)

func NewClient(cfg *config.Config) *dynamodb.Client {
	client := dynamodb.NewFromConfig(cfg.AWSConfig)
	return client
}
