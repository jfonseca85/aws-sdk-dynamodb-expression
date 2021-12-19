package app

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jfonseca85/aws-sdk-dynamodb-expression/config"
)

func NextVersion(cfg *config.Config, clientDynamoDB *dynamodb.Client, id string) string {
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(cfg.Viper.GetString(AppTable)),
		Key: map[string]types.AttributeValue{
			"id":      &types.AttributeValueMemberS{Value: id},
			"version": &types.AttributeValueMemberS{Value: AttributeVersionReservedVersion},
		},
		UpdateExpression: aws.String("SET Latest = if_not_exists(Latest, :defaultval) + :incrval"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":defaultval": &types.AttributeValueMemberN{Value: "0"},
			":incrval":    &types.AttributeValueMemberN{Value: "1"},
		},
		ReturnValues: types.ReturnValueAllNew,
	}
	output, err := clientDynamoDB.UpdateItem(context.TODO(), input)
	if err != nil {
		fmt.Println("Erro in NextVersion " + err.Error())
	}

	return fmt.Sprintf("v%s", output.Attributes["Latest"].(*types.AttributeValueMemberN).Value)

}
