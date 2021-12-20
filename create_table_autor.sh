#!/bin/bash
echo Criando a tabela Music no AWS DynamoDB
echo $(date) # Will print the output of date command
aws dynamodb create-table \
    --table-name dynamodb-table-app \
    --attribute-definitions \
        AttributeName=id,AttributeType=S \
        AttributeName=version,AttributeType=S \
    --key-schema \
        AttributeName=id,KeyType=HASH \
        AttributeName=version,KeyType=RANGE \
--provisioned-throughput \
        ReadCapacityUnits=10,WriteCapacityUnits=5
#--endpoint-url http://localhost:8000
