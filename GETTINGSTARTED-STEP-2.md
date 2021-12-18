## [ HandsOn ] - Configuração do Projeto Feature Expression do AWS DynamoDB

Para começar a trabalhar com o SDK, configure seu projeto para módulos Go, e recupere as dependências do SDK com `go get`

Este exemplo mostra como você pode usar o SDK v2 para fazer uma solicitação  de API  usando cliente [ Amazon DynamoDB] do SDK.

###### Iniciando o Projeto

```sh
$ mkdir ~/aws-sdk-dynamodb-expression
$ cd ~/aws-sdk-dynamodb-expression
$ go mod init aws-sdk-dynamodb-expression
```

###### Adicionando dependência do SDK

```sh
$ go get github.com/aws/aws-sdk-go-v2/aws
$ go get github.com/aws/aws-sdk-go-v2/config
$ go get github.com/aws/aws-sdk-go-v2/service/dynamodb
$ go get github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression
```

###### [ TDD ] - Criação dos nosso primeiro teste

No seu editor de preferência adicione o código fonte no arquivo `expression_test.go` no package expression_test

```
import (
"context"
"fmt"
"log"


"github.com/aws/aws-sdk-go-v2/aws"
"github.com/aws/aws-sdk-go-v2/config"
"github.com/aws/aws-sdk-go-v2/service/dynamodb"

)

func main() {
// Usando a configuração padrão do SDK's, carregando configuração adicionais
// e o valores de suas de suas credenciais  das variáveis de ambiente
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("sa-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	svc := dynamodb.NewFromConfig(cfg)

	tableName := "hello-world-labs-table-01"
	attributeName := "id-labs-table-01"

	var keytable []types.KeySchemaElement
	var schemaElement = types.KeySchemaElement{
		AttributeName: &attributeName,
		KeyType:       types.KeyTypeHash,
	}
	keytable = append(keytable, schemaElement)

	var attributeDefinitionList []types.AttributeDefinition
	var attributeDefinition = types.AttributeDefinition{
		AttributeName: &attributeName,
		AttributeType: "S",
	}
	attributeDefinitionList = append(attributeDefinitionList, attributeDefinition)

	createTableOutput := dynamodb.CreateTableInput{

		AttributeDefinitions: attributeDefinitionList,
		TableName:            &tableName,
		KeySchema:            keytable,
		BillingMode:          types.BillingModePayPerRequest,
	}

	_, err = svc.CreateTable(context.TODO(), &createTableOutput)

	// Build the request with its input parameters
	resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	fmt.Println("Tables:")
	for _, tableName := range resp.TableNames {
		fmt.Println(tableName)
	}

}
```

###### Compile and Execute

```sh
$ go run .
Table:
hello-world-labs-table-01
```
