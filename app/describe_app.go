package app

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jfonseca85/aws-sdk-dynamodb-expression/configlocal"
)

func GetAppParams() []*Param {
	return []*Param{
		{
			Name:     argId,
			Type:     "string",
			Required: true,
		},
		{
			Name:     argVersion,
			Type:     "string",
			Required: false,
			Default:  "latest",
		},
	}
}

func GetApp(ctx context.Context, id string, version string) (*Model, error) {
	fmt.Println("Getting GetAppBy")

	if id == "" && version == "" {
		return nil, fmt.Errorf("The id and version fields are required")

	}

	// Reserved version
	if version == AttributeVersionReservedVersion {
		return nil, fmt.Errorf("App with id %s and version %s not found", id, version)
	}

	getItemInput := buildGetItemInput(id, version)
	//clientDynamoBD := config.AWSClient.DynamoDB()

	cfg, err := configlocal.NewConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg.AWSClient)

	output, err := client.GetItem(ctx, getItemInput)
	if err != nil {
		fmt.Errorf("Erro ao buscar o App com ID: %s e VERSION: %s ", id, version)
		return nil, err
	}

	response, err := buildResponseGet(output)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}
	return response, nil

}

func buildGetItemInput(id string, version string) *dynamodb.GetItemInput {
	result := dynamodb.GetItemInput{
		TableName: aws.String(AttributeTableNameApp),
		Key: map[string]types.AttributeValue{
			"ID":      &types.AttributeValueMemberS{Value: id},
			"Version": &types.AttributeValueMemberS{Value: version},
		},
	}
	return &result
}

func buildResponseGet(out *dynamodb.GetItemOutput) (*Model, error) {
	var response Model
	err := attributevalue.UnmarshalMap(out.Item, &response)
	return &response, err
}
