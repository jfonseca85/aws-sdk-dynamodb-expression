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
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/jfonseca85/aws-sdk-dynamodb-expression/configlocal"
)

func CreateApp(args map[string]string) (string, error) {
	fmt.Println("Invoke CreateApp")

	fmt.Println("Invoke ValidateParams")
	err := ValidateParams(args, CreateAppParams())
	if err != nil {
		fmt.Println("Erro durante a validação do parametros: ", err.Error())
		return "", err
	}

	cfg, err := configlocal.NewConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg.AWSClient)
	nextVersion := NextVersion(client, args["id"])
	input := buildInput(args, nextVersion)

	output, err := Update(client, input)
	if err != nil {
		fmt.Println("Erro no Update/Create App, " + err.Error())
		return "", err
	}

	ret, err := json.Marshal(output)
	if err != nil {
		fmt.Println("Erro ao fazer o bind do App, " + err.Error())
		return "", err
	}
	return string(ret), nil

}

func Update(clientDynamoDB *dynamodb.Client, input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	output, err := clientDynamoDB.UpdateItem(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func buildInput(args map[string]string, nextVersion string) *dynamodb.UpdateItemInput {
	update := buildUpdateBuilder(args)
	// Create the DynamoDB expression from the Update.
	expr, err := expression.NewBuilder().
		WithUpdate(update).
		Build()
	if err != nil {
		fmt.Println(err)
	}

	result := &dynamodb.UpdateItemInput{
		TableName:                 aws.String(AttributeTableNameApp),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		Key: map[string]types.AttributeValue{
			"id":      &types.AttributeValueMemberS{Value: args["id"]},
			"version": &types.AttributeValueMemberS{Value: nextVersion},
		},
		//UpdateExpression: aws.String("SET Latest = if_not_exists(Latest, :defaultval) + :incrval, #ArnAsf= :asf_value, #Status = :status_value, #Document = :document_value"),
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
	//UpdateExpression: aws.String("SET Latest = if_not_exists(Latest, :defaultval) + :incrval,
	//update = update.Set(expression.Name("Latest"), expression.IfNotExists(expression.Name("Latest"), expression.Value(0)))
	//update = update.Set(expression.Name("Latest"), expression.Plus(expression.Name("Latest"), expression.IfNotExists(expression.Name("Latest"), expression.Value(0)), expression.Value(1)))
	return update
}
