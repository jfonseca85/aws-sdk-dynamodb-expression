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
		"version": "v3",
	}
	expect := "v3"

	//Ação
	result, err := app.GetApp(ctx, args)

	//Validação
	if err != nil {
		t.Errorf("Erro ao obter oa App: %v/n", err.Error())
	}

	if expect != result.Version {
		t.Errorf("A versão retornada deve ser igual a %v/n porém retornou version: %v/n", expect, result.Version)
	}
	fmt.Println("Apps retornados: ", result)

}

func Test_GetApp_Missing_Version(t *testing.T) {
	//Cenário
	ctx := context.TODO()
	args := map[string]string{
		"id": "1",
	}

	//Ação
	result, err := app.GetApp(ctx, args)

	//Validação
	if err != nil {
		t.Errorf("Erro ao obter oa App: %v/n", err.Error())
	}
	if result == nil {
		t.Errorf("Devera retorna um objeto")
	}
	fmt.Println("App retornado: ", result)

}

func Test_GetApp_Empty_Version(t *testing.T) {
	//Cenário
	ctx := context.TODO()
	args := map[string]string{
		"id":      "1",
		"version": "",
	}

	//Ação
	result, err := app.GetApp(ctx, args)

	//Validação
	if err == nil {
		t.Errorf("Deverá retornar error, versão não esta vazia")
	}
	fmt.Println("App retornado: ", result)

}

func Test_GetApp_Incorret_Version(t *testing.T) {
	//Cenário
	ctx := context.TODO()
	args := map[string]string{
		"id":      "1",
		"version": "gghgajsgj",
	}

	//Ação
	result, err := app.GetApp(ctx, args)

	//Validação
	if err != nil {
		t.Errorf("Erro ao buscar o App %v/n", err.Error())
	}

	if result.ID != "" {
		t.Errorf("Deverá retornar vazia versão inexistente e fora do padrão")
	}
	fmt.Println("App retornado: ", result)

}
