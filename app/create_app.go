package app

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jfonseca85/aws-sdk-dynamodb-expression/config"
)

func CreateApp(ctx context.Context, cfg *config.Config, args map[string]string) (string, error) {
	fmt.Println("Invoke CreateApp")
	err := ValidateParams(args, CreateAppParams())
	if err != nil {
		return "", err
	}

	result, _ := UpdateAppMiddleware(ctx, cfg, args)
	ret, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(ret), nil
}

func CreateAppParams() []*Param {

	return []*Param{
		{
			Name:     argId,
			Type:     "string",
			Required: true,
		},
		{
			Name:     argDocument,
			Type:     "yaml",
			Required: true,
		},
	}
}

func UpdateAppMiddleware(ctx context.Context, cfg *config.Config, args map[string]string) (*Model, error) {

	client := NewClient(cfg)

	nextVersion := NextVersion(cfg, client, args["id"])
	input := buildInput(cfg, args, nextVersion)

	output, err := updateItem(client, input)
	if err != nil {
		return nil, err
	}

	ret, err := buildResponseCreate(output)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func buildResponseCreate(out *dynamodb.UpdateItemOutput) (*Model, error) {
	var response Model
	err := attributevalue.UnmarshalMap(out.Attributes, &response)

	return &response, err
}

func updateItem(clientDynamoDB *dynamodb.Client, input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	output, err := clientDynamoDB.UpdateItem(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func buildInput(cfg *config.Config, args map[string]string, nextVersion string) *dynamodb.UpdateItemInput {
	update := buildUpdateBuilder(args)
	expr, err := expression.NewBuilder().
		WithUpdate(update).
		Build()
	if err != nil {
		fmt.Println(err)
	}

	result := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(cfg.Viper.GetString(AppTable)),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		Key: map[string]types.AttributeValue{
			"id":      &types.AttributeValueMemberS{Value: args["id"]},
			"version": &types.AttributeValueMemberS{Value: nextVersion},
		},
		UpdateExpression: expr.Update(),
		ReturnValues:     types.ReturnValueAllNew,
	}
	return result
}

func buildUpdateBuilder(expressions map[string]string) expression.UpdateBuilder {
	var update expression.UpdateBuilder
	for k, v := range expressions {
		if k == "id" || k == "version" {
			continue
		}
		update = update.Set(
			expression.Name(k),
			expression.Value(v),
		)
	}
	return update
}
