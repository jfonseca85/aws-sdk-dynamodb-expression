//Esse adaptador de driver deve ser capaz de transformar uma solicitação http em uma chamada de serviço.
//Ter todos os componentes separados uns dos outros nos dá a vantagem de implementá-los e
//Testá-los isoladamente ou podemos facilmente paralelizar o trabalho com a ajuda de outros membros da equipe.
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

func ListApps(ctx context.Context, cfg *config.Config, args map[string]string) (string, error) {
	fmt.Println("Invoke ListApps")

	err := ValidateParams(args, ListAppsParams())
	if err != nil {
		return "", err
	}

	// Using the Config value, create the DynamoDB client
	client := NewClient(cfg)

	out, err := listApps(client, buildScanInput(cfg))
	if err != nil {
		fmt.Println("Unable to list apps, " + err.Error())
		return "", err
	}
	return buildResponse(out)
}

func ListAppsParams() []*Param {
	return []*Param{}
}

func buildScanInput(cfg *config.Config) *dynamodb.ScanInput {
	input := &dynamodb.ScanInput{
		TableName:        aws.String(cfg.Viper.GetString(AppTable)),
		FilterExpression: aws.String("#version <> :version"),
		ExpressionAttributeNames: map[string]string{
			"#version": "version",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":version": &types.AttributeValueMemberS{Value: AttributeVersionReservedVersion},
		},
	}

	return input
}

func listApps(clientDynamoDB *dynamodb.Client, params *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	output, err := clientDynamoDB.Scan(context.TODO(), params)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func buildResponse(scanOutput *dynamodb.ScanOutput) (string, error) {

	var response []Model
	err := attributevalue.UnmarshalListOfMaps(scanOutput.Items, &response)
	if err != nil {
		return "", err
	}

	ret, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	return string(ret), nil
}
