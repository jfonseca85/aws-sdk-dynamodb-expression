## Getting started

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



###### Criar uma tabela Music

Neste exemplo utilizaresmo a tabela Music para fazer os testes de como usar a feature exepression do DynamoDB

```sh
aws dynamodb create-table \
    --table-name Music \
    --attribute-definitions \
        AttributeName=Artist,AttributeType=S \
        AttributeName=SongTitle,AttributeType=S \
    --key-schema \
        AttributeName=Artist,KeyType=HASH \
        AttributeName=SongTitle,KeyType=RANGE \
--provisioned-throughput \
        ReadCapacityUnits=10,WriteCapacityUnits=5
```

O uso de create-table retorna o seguinte resultado de exemplo.

```sh
{
    "TableDescription": {
        "TableArn": "arn:aws:dynamodb:us-west-2:522194210714:table/Music",
        "AttributeDefinitions": [
            {
                "AttributeName": "Artist",
                "AttributeType": "S"
            },
            {
                "AttributeName": "SongTitle",
                "AttributeType": "S"
            }
        ],
        "ProvisionedThroughput": {
            "NumberOfDecreasesToday": 0,
            "WriteCapacityUnits": 5,
            "ReadCapacityUnits": 10
        },
        "TableSizeBytes": 0,
        "TableName": "Music",
        "TableStatus": "CREATING", 
        "TableId": "d04c7240-0e46-435d-b231-d54091fe1017",
        "KeySchema": [
            {
                "KeyType": "HASH",
                "AttributeName": "Artist"
            },
            {
                "KeyType": "RANGE",
                "AttributeName": "SongTitle"
            }
        ],
        "ItemCount": 0,
        "CreationDateTime": 1558028402.69
    }
}
```

Para verificar se o DynamoDB terminou de criar a tabela Music, use o comando describe-table.
```sh
 aws dynamodb describe-table --table-name Music | grep TableStatus
 
```

Esse comando retorna o seguinte resultado. Quando o DynamoDB conclui a criação da tabela, o valor do campo TableStatus é definido como ACTIVE.

```sh
"TableStatus": "ACTIVE",
```

###### Escrevendo o Código



No seu editor de preferência adicione o código fonte no arquivo `main.go` package main

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
