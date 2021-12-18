package expression_test

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/jfonseca85/aws-sdk-dynamodb-expression/configlocal"
)

// Using Update Expression
//
// This example updates an item in the Music table. It adds a new attribute (Year) and
// modifies the AlbumTitle attribute.  All of the attributes in the item, as they appear
// after the update, are returned in the response.

func ExampleBuilder_WithUpdate() {

	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := configlocal.NewConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg.AWSClient)

	//client := dynamodb.NewFromConfig(cfg)

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
