package app_test

import (
	"fmt"
	"testing"

	"github.com/jfonseca85/aws-sdk-dynamodb-expression/app"
)

func Test_CreateApp_Sucess(t *testing.T) {
	//Cenário: Criar args com todos os campos obrigatórios (id e document)
	args := map[string]string{
		"id":       "1",
		"document": "QVdTVGVtcGxhdGVGb3JtYXRWZXJzaW9uOiAnMjAxMC0wOS0wOScKUmVzb3VyY2VzOgogIG15U3RhY2s6CiAgICBUeXBlOiBBV1M6OkNsb3VkRm9ybWF0aW9uOjpTdGFjawogICAgUHJvcGVydGllczoKICAgICAgVGVtcGxhdGVVUkw6IGh0dHBzOi8vczMuYW1hem9uYXdzLmNvbS9jbG91ZGZvcm1hdGlvbi10ZW1wbGF0ZXMtdXMtZWFzdC0xL1MzX0J1Y2tldC50ZW1wbGF0ZQogICAgICBUaW1lb3V0SW5NaW51dGVzOiAnNjAnCk91dHB1dHM6CiAgU3RhY2tSZWY6CiAgICBWYWx1ZTogIVJlZiBteVN0YWNrCiAgT3V0cHV0RnJvbU5lc3RlZFN0YWNrOgogICAgVmFsdWU6ICFHZXRBdHQgbXlTdGFjay5PdXRwdXRzLkJ1Y2tldE5hbWU=",
	}

	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}

	//Ação: Cria o App

	result, err := app.CreateApp(ctx, cfg, args)

	//Validação: Valida se criou com sucesso
	if err != nil {
		t.Errorf("Erro durante Criação do App: %v/n ", err.Error())
	}
	fmt.Println("App criado com sucesso: ", result)

}

func Test_CreateApp_Sending_Version(t *testing.T) {
	//Cenário: Cria args inserindo a version não usada para criar o App.
	args := map[string]string{
		"id":       "1",
		"version":  "v1",
		"cluster":  "itau-service-mesh",
		"document": "QVdTVGVtcGxhdGVGb3JtYXRWZXJzaW9uOiAnMjAxMC0wOS0wOScKUmVzb3VyY2VzOgogIG15U3RhY2s6CiAgICBUeXBlOiBBV1M6OkNsb3VkRm9ybWF0aW9uOjpTdGFjawogICAgUHJvcGVydGllczoKICAgICAgVGVtcGxhdGVVUkw6IGh0dHBzOi8vczMuYW1hem9uYXdzLmNvbS9jbG91ZGZvcm1hdGlvbi10ZW1wbGF0ZXMtdXMtZWFzdC0xL1MzX0J1Y2tldC50ZW1wbGF0ZQogICAgICBUaW1lb3V0SW5NaW51dGVzOiAnNjAnCk91dHB1dHM6CiAgU3RhY2tSZWY6CiAgICBWYWx1ZTogIVJlZiBteVN0YWNrCiAgT3V0cHV0RnJvbU5lc3RlZFN0YWNrOgogICAgVmFsdWU6ICFHZXRBdHQgbXlTdGFjay5PdXRwdXRzLkJ1Y2tldE5hbWU=",
	}

	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}

	//Ação: Cria o App e pega a ultima versão descartando a versão passada
	result, err := app.CreateApp(ctx, cfg, args)

	//Validação: Valida se foi criado corretamente utilizando a ultima versão
	if err != nil {
		t.Errorf("Erro durante Criação do App: %v/n ", err.Error())
	}
	fmt.Println("App criado com sucesso: ", result)
}

func Test_CreateApp_Empty_Args(t *testing.T) {
	//Cenário: Cria args vazia
	args := map[string]string{}

	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}
	//Ação
	_, err = app.CreateApp(ctx, cfg, args)

	//Validação: Valida se retornou erro, faltam os campos obrigatórios (id e document)
	if err == nil {
		t.Errorf("A CreateApp deverá retornar erro (Campo Id é ogrigatório)")
	}
}

func Test_CreateApp_Null_Args(t *testing.T) {
	//Cenário:
	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}
	//Ação: Chama cria App passando null
	_, err = app.CreateApp(ctx, cfg, nil)

	//Validação:  Valida se retornou erro, faltam os campos obrigatórios (id e document)
	if err == nil {
		t.Errorf("A CreateApp deverá retornar erro (Campo Id é ogrigatório)")
	}
}
