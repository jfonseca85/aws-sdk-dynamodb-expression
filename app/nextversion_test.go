package app_test

import (
	"fmt"
	"testing"

	"github.com/jfonseca85/aws-sdk-dynamodb-expression/app"
)

func Test_NextVersion(t *testing.T) {

	//A Nova Versão deve ser incrementada em 1 (Ex.: Versão atual: v3, Nova Versão: v4)

	//Cenario: Cria o id do App
	client := app.InitClient()
	idTable := "1"
	//Ação: Obtem nova Versão
	nextVersion := app.NextVersion(client, idTable)
	fmt.Printf("Exibindo a nova versão do App: %q/n", nextVersion)

	//Validação:
	//TODO: (Jorge Luis) Fazer um GET pegando a última versão depois executa o NextVersion

}
