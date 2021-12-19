package app_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jfonseca85/aws-sdk-dynamodb-expression/app"
)

/*
	Obtem o App pasando o ID e Version"
*/
func Test_GetApp_Sucess(t *testing.T) {
	//Cenário
	ctx := context.TODO()
	args := map[string]string{
		"id":      "1",
		"version": "v4",
	}

	//Ação
	resultApp, err := app.GetApp(ctx, args)
	if err != nil {
		t.Errorf("Erro ao obter oa App:%v/n", err.Error())
	}
	//Validação
	fmt.Println("Apps retornados: ", resultApp)

}
