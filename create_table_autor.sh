#!/bin/bash
echo Criando a tabela Music no AWS DynamoDB
echo $(date) # Will print the output of date command
aws dynamodb create-table \
    --table-name dynamodb-table-app \
    --attribute-definitions \
        AttributeName=ID,AttributeType=S \
        AttributeName=Version,AttributeType=S \
    --key-schema \
        AttributeName=ID,KeyType=HASH \
        AttributeName=Version,KeyType=RANGE \
--provisioned-throughput \
        ReadCapacityUnits=10,WriteCapacityUnits=5