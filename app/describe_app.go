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
		return nil, err
	}

	cfg, err := configlocal.NewConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	client := dynamodb.NewFromConfig(cfg.AWSClient)

	getItemInput, err := buildGetItemInput(args["id"], args["version"])
	if err != nil {
		return nil, err
	}

	output, err := client.GetItem(ctx, getItemInput)
	if err != nil {
		return nil, err
	}

	response, err := buildResponseGet(output)
	if err != nil {
		return nil, err
	}
	return response, nil

}

func buildGetItemInput(id string, version string) (*dynamodb.GetItemInput, error) {
	versionInput := id
	if version == AttributeVersionLatestVersion {
		lastedVersion, err := getLastedVersion(id)
		if err != nil {
			return nil, err
		}
		versionInput = lastedVersion
	}

	return &dynamodb.GetItemInput{
		TableName: aws.String(AttributeTableNameApp),
		Key: map[string]types.AttributeValue{
			"id":      &types.AttributeValueMemberS{Value: id},
			"version": &types.AttributeValueMemberS{Value: versionInput},
		},
	}, nil
}

func buildResponseGet(out *dynamodb.GetItemOutput) (*Model, error) {
	var response Model
	err := attributevalue.UnmarshalMap(out.Item, &response)
	return &response, err
}

func getLastedVersion(id string) (string, error) {
	cfg, err := configlocal.NewConfig(context.TODO())
	if err != nil {
		return "", err
	}
	client := dynamodb.NewFromConfig(cfg.AWSClient)

	input := dynamodb.GetItemInput{
		TableName: aws.String(AttributeTableNameApp),
		Key: map[string]types.AttributeValue{
			"id":      &types.AttributeValueMemberS{Value: id},
			"version": &types.AttributeValueMemberS{Value: AttributeVersionReservedVersion},
		},
	}
	result, err := client.GetItem(context.TODO(), &input)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("v%s", result.Item["Latest"].(*types.AttributeValueMemberN).Value), nil
}
