package app

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jfonseca85/aws-sdk-dynamodb-expression/configlocal"
)

func NewClient(cfg *configlocal.Viperloadconfig) *dynamodb.Client {

	//cfg, err := cfg.NewConfig(context.TODO())
	//if err != nil {
	//	fmt.Println("unable to load SDK config:,", err.Error())
	//}
	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg.AWSClient)

	return client
}

//Utilizado nos testes
func InitClient() (context.Context, *configlocal.Viperloadconfig, error) {
	ctx := context.TODO()
	cfg, err := configlocal.NewConfig(ctx)
	if err != nil {
		return nil, nil, err
	}
	return ctx, cfg, nil
}
