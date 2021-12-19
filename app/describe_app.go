package app

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jfonseca85/aws-sdk-dynamodb-expression/config"
)

func GetApp(ctx context.Context, cfg *config.Config, args map[string]string) (string, error) {
	fmt.Println("Invoke GetApp")

	err := ValidateParams(args, GetAppParams())
	if err != nil {
		return "", err
	}

	client := NewClient(cfg)

	getItemInput, err := buildGetItemInput(ctx, cfg, args["id"], args["version"])
	if err != nil {
		return "", err
	}

	output, err := client.GetItem(ctx, getItemInput)
	if err != nil {
		return "", err
	}

	result, err := buildResponseGet(output)
	if err != nil {
		return "", err
	}

	ret, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(ret), nil

}

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

func buildGetItemInput(ctx context.Context, cfg *config.Config, id string, version string) (*dynamodb.GetItemInput, error) {
	versionInput := version
	if version == AttributeVersionLatestVersion {
		lastedVersion, err := getLastedVersion(ctx, cfg, id)
		if err != nil {
			return nil, err
		}
		versionInput = lastedVersion
	}

	return &dynamodb.GetItemInput{
		TableName: aws.String(cfg.Viper.GetString(AppTable)),
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

func getLastedVersion(ctx context.Context, cfg *config.Config, id string) (string, error) {
	client := NewClient(cfg)
	input := dynamodb.GetItemInput{
		TableName: aws.String(cfg.Viper.GetString(AppTable)),
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
