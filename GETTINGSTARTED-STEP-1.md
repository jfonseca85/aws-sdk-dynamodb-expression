## Getting started

Neste passo mostraremos como criar a tabela Muisc na sua conta da AWS.

###### Criar uma tabela Music

Neste exemplo utilizaresmo a tabela Music para fazer os testes de como usar a feature exepression do DynamoDB

Execute o script create_table.sh para cria a tabela Music no DynamoDB

```sh
./create_table.sh
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
