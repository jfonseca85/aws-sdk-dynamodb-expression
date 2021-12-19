package app_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jfonseca85/aws-sdk-dynamodb-expression/app"
)

func Test_listApp_Sucess(t *testing.T) {
	//cenário
	args := map[string]string{}
	ctx := context.TODO()

	//Ação
	listApps, err := app.ListApps(ctx, args)
	if err != nil {
		fmt.Println("Erro ao listar os Apps:", err.Error())

	}
	//Validação
	fmt.Println("Apps retornados: ", listApps)
}
