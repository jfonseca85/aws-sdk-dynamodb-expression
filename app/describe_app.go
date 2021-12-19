package app

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jfonseca85/aws-sdk-dynamodb-expression/configlocal"
)

func GetApp(ctx context.Context, args map[string]string) (*Model, error) {

	fmt.Println("Invoke GetApp")

	err := ValidateParams(args, GetAppParams())
	if err != nil {
		fmt.Println("Erro a fazer a validação:", err.Error())
		return nil, err
	}

	getItemInput := buildGetItemInput(args["id"], args["version"])
	//clientDynamoBD := config.AWSClient.DynamoDB()

	cfg, err := configlocal.NewConfig(context.TODO())
	if err != nil {
		fmt.Println("Erro ao carregar a config local", err.Error())
		return nil, err
	}
	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg.AWSClient)

	output, err := client.GetItem(ctx, getItemInput)
	if err != nil {
		return nil, err
	}

	response, err := buildResponseGet(output)
	if err != nil {
		fmt.Println("Erro ao fazer o bind de resposta", err.Error())
		return nil, err
	}
	return response, nil

}

func buildGetItemInput(id string, version string) *dynamodb.GetItemInput {
	result := dynamodb.GetItemInput{
		TableName: aws.String(AttributeTableNameApp),
		Key: map[string]types.AttributeValue{
			"id":      &types.AttributeValueMemberS{Value: id},
			"version": &types.AttributeValueMemberS{Value: version},
		},
	}
	return &result
}

func buildResponseGet(out *dynamodb.GetItemOutput) (*Model, error) {
	var response Model
	err := attributevalue.UnmarshalMap(out.Item, &response)
	return &response, err
}
