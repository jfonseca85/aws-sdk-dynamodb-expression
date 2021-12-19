package app

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jfonseca85/aws-sdk-dynamodb-expression/configlocal"
)

func InitClient() *dynamodb.Client {

	cfg, err := configlocal.NewConfig(context.TODO())
	if err != nil {
		fmt.Println("unable to load SDK config:,", err.Error())
	}
	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg.AWSClient)

	return client
}
