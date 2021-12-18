//Esse adaptador de driver deve ser capaz de transformar uma solicitação http em uma chamada de serviço.
//Ter todos os componentes separados uns dos outros nos dá a vantagem de implementá-los e
//Testá-los isoladamente ou podemos facilmente paralelizar o trabalho com a ajuda de outros membros da equipe.
package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jfonseca85/aws-sdk-dynamodb-expression/configlocal"
)

func CreateAppParams() []*Param {

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
		{
			Name:     argDocument,
			Type:     "yaml",
			Required: true,
		},
	}
}

func CreateApp(args map[string]string) (string, error) {
	log.Printf("Getting  CreateApp>>> ")

	err := ValidateParams(args, CreateAppParams())

	if err != nil {
		return "", err
	}

	body := &Model{
		ID:       args[argId],
		Version:  args[argVersion],
		Document: args[argDocument],
	}

	//client := config.AWSClient.DynamoDB()
	cfg, err := configlocal.NewConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg.AWSClient)
	NextVersion(client, body)

	output, err := Update(client, body)
	if err != nil {
		fmt.Println("Erro no Update/Create App, " + err.Error())
		return "", err
	}

	ret, err := json.Marshal(output)
	if err != nil {
		fmt.Println("Erro ao fazer o bind do response , " + err.Error())
		return "", err
	}
	return string(ret), nil

}

func Update(clientDynamoDB *dynamodb.Client, model *Model) (*dynamodb.UpdateItemOutput, error) {
	log.Printf("Chamando o app.Update>>> ")

	input := buildInput(model)

	output, err := clientDynamoDB.UpdateItem(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func buildInput(model *Model) *dynamodb.UpdateItemInput {
	result := &dynamodb.UpdateItemInput{
		TableName: aws.String(AttributeTableNameApp),
		Key: map[string]types.AttributeValue{
			"ID":      &types.AttributeValueMemberS{Value: model.ID},
			"Version": &types.AttributeValueMemberS{Value: model.Version},
		},
		UpdateExpression: aws.String("SET Latest = if_not_exists(Latest, :defaultval) + :incrval, #ArnAsf= :asf_value, #Status = :status_value, #Document = :document_value"),
		ExpressionAttributeNames: map[string]string{
			"#ArnAsf":   "ArnAsf",
			"#Status":   "Status",
			"#Document": "Document",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":defaultval":     &types.AttributeValueMemberN{Value: "0"},
			":incrval":        &types.AttributeValueMemberN{Value: "1"},
			":asf_value":      &types.AttributeValueMemberS{Value: model.ArnAsf},
			":status_value":   &types.AttributeValueMemberS{Value: model.Status},
			":document_value": &types.AttributeValueMemberS{Value: model.Document},
		},
		ReturnValues: types.ReturnValueAllNew,
	}
	return result
}
