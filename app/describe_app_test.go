package app_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jfonseca85/aws-sdk-dynamodb-expression/app"
)

func Test_GetApp_Sucess(t *testing.T) {
	//Cenário: Cria args para obter o App (id: 1 e version: v3)
	args := map[string]string{
		"id":      "1",
		"version": "v3",
	}
	expect := "v3"

	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}

	//Ação: Obtem o App
	result, err := app.GetApp(ctx, cfg, args)

	byt := []byte(result)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	strs := dat["version"].(interface{})
	str1 := strs.(string)

	//Validação: Valida o retorno do App (id: 1 e version: v3)
	if err != nil {
		t.Errorf("Erro ao obter oa App: %v/n", err.Error())
	}

	if expect != str1 {
		t.Errorf("A versão retornada deve ser igual a %v/n porém retornou version: %v/n", expect, str1)
	}
	fmt.Println("Apps retornados: ", result)

}

func Test_GetApp_Missing_Version(t *testing.T) {
	//Cenário: Cria args faltando version ( obrigatório )
	args := map[string]string{
		"id": "1",
	}
	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}

	//Ação: Obtém o App
	result, err := app.GetApp(ctx, cfg, args)

	//Validação: Deverá retornar a última versão do App (id: 1)
	if err != nil {
		t.Errorf("Erro ao obter oa App: %v/n", err.Error())
	}
	if result == "" {
		t.Errorf("Devera retorna um objeto")
	}
	fmt.Println("App retornado: ", result)

}

func Test_GetApp_Empty_Version(t *testing.T) {
	//Cenário: Cria args com version vazia ( INCORRETO )
	args := map[string]string{
		"id":      "1",
		"version": "",
	}

	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}

	//Ação: Obtem o App
	result, err := app.GetApp(ctx, cfg, args)

	//Validação: Deverá retornar error, versão esta vazia
	if err == nil {
		t.Errorf("Deverá retornar error, versão esta vazia")
	}
	fmt.Println("App retornado: ", result)

}

func Test_GetApp_Incorret_Version(t *testing.T) {
	//Cenário: Cria args com versão inexistente

	args := map[string]string{
		"id":      "1",
		"version": "gghgajsgj",
	}

	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}

	//Ação: Obtém App
	result, err := app.GetApp(ctx, cfg, args)

	byt := []byte(result)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	strs := dat["id"].(interface{})
	str1 := strs.(string)

	//Validação: Deverá retornar vazia versão inexistente
	if err != nil {
		t.Errorf("Erro ao buscar o App %v/n", err.Error())
	}

	if str1 != "" {
		t.Errorf("Deverá retornar vazia versão inexistente e fora do padrão")
	}
	fmt.Println("App retornado: ", result)

}
