package app_test

import (
	"fmt"
	"testing"

	"github.com/jfonseca85/aws-sdk-dynamodb-expression/app"
)

func Test_listApp_Sucess(t *testing.T) {
	//Cenário: Cria args vazia ( Estes args não são usados na consulta de App)
	args := map[string]string{}
	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}

	//Ação: Lista Apps
	listApps, err := app.ListApps(ctx, cfg, args)
	if err != nil {
		fmt.Println("Erro ao listar os Apps:", err.Error())
	}
	//Validação: Retorna a lista de Apps
	fmt.Println("Apps retornados: ", listApps)
}

func Test_listApp_Params_Null(t *testing.T) {
	//Cenário: Cria args Null ( Estes args não são usados na consulta de App)
	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}
	//Ação: Lista Apps
	listApps, err := app.ListApps(ctx, cfg, nil)
	if err != nil {
		fmt.Println("Erro ao listar os Apps:", err.Error())
	}
	//Validação: Retorna a lista de Apps
	fmt.Println("Apps retornados: ", listApps)
}

func Test_listApp_With_Params(t *testing.T) {
	//Cenário: Cria args com params ( Estes args não são usados na consulta de App)
	args := map[string]string{
		"id":      "1",
		"version": "v2",
	}
	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}

	//Ação: Lista Apps
	listApps, err := app.ListApps(ctx, cfg, args)
	if err != nil {
		fmt.Println("Erro ao listar os Apps:", err.Error())
	}
	//Validação: Retorna a lista de Apps
	fmt.Println("Apps retornados: ", listApps)
}
