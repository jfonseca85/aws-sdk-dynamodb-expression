package app_test

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var config = struct {
	LoadDefaultConfig func(ctx context.Context) (aws.Config, error)
}{
	LoadDefaultConfig: func(ctx context.Context) (aws.Config, error) {
		return aws.Config{}, nil
	},
}

func ExampleBuilder_WithUpdate() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client := dynamodb.NewFromConfig(cfg)

	// Create an update to set two fields in the table.
	update := expression.Set(
		expression.Name("Year"),
		expression.Value(2015),
	).Set(
		expression.Name("AlbumTitle"),
		expression.Value("Louder Than Ever"),
	)

	// Create the DynamoDB expression from the Update.
	expr, err := expression.NewBuilder().
		WithUpdate(update).
		Build()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Use the built expression to populate the DynamoDB UpdateItem API
	// input parameters.
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		Key: map[string]types.AttributeValue{
			"Artist":    &types.AttributeValueMemberS{Value: "Acme Band"},
			"SongTitle": &types.AttributeValueMemberS{Value: "Happy Day"},
		},
		ReturnValues:     "ALL_NEW",
		TableName:        aws.String("Music"),
		UpdateExpression: expr.Update(),
	}

	result, err := client.UpdateItem(context.TODO(), input)
	if err != nil {
		if apiErr := new(types.ProvisionedThroughputExceededException); errors.As(err, &apiErr) {
			fmt.Println("throughput exceeded")
		} else if apiErr := new(types.ResourceNotFoundException); errors.As(err, &apiErr) {
			fmt.Println("resource not found")
		} else if apiErr := new(types.InternalServerError); errors.As(err, &apiErr) {
			fmt.Println("internal server error")
		} else {
			fmt.Println(err)
		}
		return
	}

	fmt.Println(result)
}
