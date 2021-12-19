package app_test

import (
	"fmt"
	"testing"

	"github.com/jfonseca85/aws-sdk-dynamodb-expression/app"
)

func Test_UpdateApp_Sucess(t *testing.T) {
	//Cenário: Criar args com todos os campos obrigatórios para o Update (id e version )
	args := map[string]string{
		"id":       "1",
		"version":  "v2",
		"xpto":     "value_xpto",
		"document": "QVdTVGVtcGxhdGVGb3JtYXRWZXJzaW9uOiAnMjAxMC0wOS0wOScKUmVzb3VyY2VzOgogIG15U3RhY2s6CiAgICBUeXBlOiBBV1M6OkNsb3VkRm9ybWF0aW9uOjpTdGFjawogICAgUHJvcGVydGllczoKICAgICAgVGVtcGxhdGVVUkw6IGh0dHBzOi8vczMuYW1hem9uYXdzLmNvbS9jbG91ZGZvcm1hdGlvbi10ZW1wbGF0ZXMtdXMtZWFzdC0xL1MzX0J1Y2tldC50ZW1wbGF0ZQogICAgICBUaW1lb3V0SW5NaW51dGVzOiAnNjAnCk91dHB1dHM6CiAgU3RhY2tSZWY6CiAgICBWYWx1ZTogIVJlZiBteVN0YWNrCiAgT3V0cHV0RnJvbU5lc3RlZFN0YWNrOgogICAgVmFsdWU6ICFHZXRBdHQgbXlTdGFjay5PdXRwdXRzLkJ1Y2tldE5hbWU=",
	}
	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}
	//Ação: Update App
	result, err := app.UpdateApp(ctx, cfg, args)

	//Validação: Valida se atualização ocorreu com sucesso
	if err != nil {
		t.Errorf("Erro durante Criação do App: %v/n ", err.Error())
	}
	fmt.Println("App atualizado com sucesso: ", result)

}

func Test_UpdateApp_Version_Missing(t *testing.T) {
	//Cenário: Criar args faltando o campo version (obrigatório)
	args := map[string]string{
		"id":       "1",
		"xpto":     "value_xpto",
		"document": "QVdTVGVtcGxhdGVGb3JtYXRWZXJzaW9uOiAnMjAxMC0wOS0wOScKUmVzb3VyY2VzOgogIG15U3RhY2s6CiAgICBUeXBlOiBBV1M6OkNsb3VkRm9ybWF0aW9uOjpTdGFjawogICAgUHJvcGVydGllczoKICAgICAgVGVtcGxhdGVVUkw6IGh0dHBzOi8vczMuYW1hem9uYXdzLmNvbS9jbG91ZGZvcm1hdGlvbi10ZW1wbGF0ZXMtdXMtZWFzdC0xL1MzX0J1Y2tldC50ZW1wbGF0ZQogICAgICBUaW1lb3V0SW5NaW51dGVzOiAnNjAnCk91dHB1dHM6CiAgU3RhY2tSZWY6CiAgICBWYWx1ZTogIVJlZiBteVN0YWNrCiAgT3V0cHV0RnJvbU5lc3RlZFN0YWNrOgogICAgVmFsdWU6ICFHZXRBdHQgbXlTdGFjay5PdXRwdXRzLkJ1Y2tldE5hbWU=",
	}
	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}
	//Ação: Realiza o update do App
	result, err := app.UpdateApp(ctx, cfg, args)

	//Validação: Deverá retornar erro, campo version é obrigatório
	if err == nil {
		t.Errorf("Deverá retornar erro, campo version é obrigatorio")
	}
	fmt.Println("App atualizado com sucesso: ", result)

}

func Test_UpdateApp_Id_Missing(t *testing.T) {
	//Cenário: Criar args faltando o campo id (obrigatório)
	args := map[string]string{
		"version":  "v2",
		"xpto":     "value_xpto",
		"document": "QVdTVGVtcGxhdGVGb3JtYXRWZXJzaW9uOiAnMjAxMC0wOS0wOScKUmVzb3VyY2VzOgogIG15U3RhY2s6CiAgICBUeXBlOiBBV1M6OkNsb3VkRm9ybWF0aW9uOjpTdGFjawogICAgUHJvcGVydGllczoKICAgICAgVGVtcGxhdGVVUkw6IGh0dHBzOi8vczMuYW1hem9uYXdzLmNvbS9jbG91ZGZvcm1hdGlvbi10ZW1wbGF0ZXMtdXMtZWFzdC0xL1MzX0J1Y2tldC50ZW1wbGF0ZQogICAgICBUaW1lb3V0SW5NaW51dGVzOiAnNjAnCk91dHB1dHM6CiAgU3RhY2tSZWY6CiAgICBWYWx1ZTogIVJlZiBteVN0YWNrCiAgT3V0cHV0RnJvbU5lc3RlZFN0YWNrOgogICAgVmFsdWU6ICFHZXRBdHQgbXlTdGFjay5PdXRwdXRzLkJ1Y2tldE5hbWU=",
	}
	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}
	//Ação: Realiza o update do App
	result, err := app.UpdateApp(ctx, cfg, args)

	//Validação: Deverá retornar erro, campo id é obrigatório
	if err == nil {
		t.Errorf("Deverá retornar erro, campo id é obrigatorio")
	}
	fmt.Println("App atualizado com sucesso: ", result)

}

func Test_UpdateApp_Id_And_Version_Missing(t *testing.T) {
	//Cenário: Criar args faltando os campos id  e version são obrigatórios
	args := map[string]string{
		"xpto":     "value_xpto",
		"document": "QVdTVGVtcGxhdGVGb3JtYXRWZXJzaW9uOiAnMjAxMC0wOS0wOScKUmVzb3VyY2VzOgogIG15U3RhY2s6CiAgICBUeXBlOiBBV1M6OkNsb3VkRm9ybWF0aW9uOjpTdGFjawogICAgUHJvcGVydGllczoKICAgICAgVGVtcGxhdGVVUkw6IGh0dHBzOi8vczMuYW1hem9uYXdzLmNvbS9jbG91ZGZvcm1hdGlvbi10ZW1wbGF0ZXMtdXMtZWFzdC0xL1MzX0J1Y2tldC50ZW1wbGF0ZQogICAgICBUaW1lb3V0SW5NaW51dGVzOiAnNjAnCk91dHB1dHM6CiAgU3RhY2tSZWY6CiAgICBWYWx1ZTogIVJlZiBteVN0YWNrCiAgT3V0cHV0RnJvbU5lc3RlZFN0YWNrOgogICAgVmFsdWU6ICFHZXRBdHQgbXlTdGFjay5PdXRwdXRzLkJ1Y2tldE5hbWU=",
	}
	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}
	//Ação: Realiza o update do App
	result, err := app.UpdateApp(ctx, cfg, args)

	//Validação: Deverá retornar erro, campos id  e version são obrigatórios
	if err == nil {
		t.Errorf("Deverá retornar erro, campos id  e version são obrigatórios")
	}
	fmt.Println("App atualizado com sucesso: ", result)

}

func Test_UpdateApp_Params_Empty(t *testing.T) {
	//Cenário: Criar args faltando os campos id  e version são obrigatórios
	args := map[string]string{}
	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}
	//Ação: Realiza o update do App
	result, err := app.UpdateApp(ctx, cfg, args)

	//Validação: Deverá retornar erro, campos id  e version são obrigatórios
	if err == nil {
		t.Errorf("Deverá retornar erro, campos id  e version são obrigatórios")
	}
	fmt.Println("App atualizado com sucesso: ", result)

}

func Test_UpdateApp_Params_Null(t *testing.T) {
	//Cenário: Criar args faltando os campos id  e version são obrigatórios
	ctx, cfg, err := NewClientLocal()
	if err != nil {
		return
	}

	//Ação: Realiza o update do App
	result, err := app.UpdateApp(ctx, cfg, nil)

	//Validação: Deverá retornar erro, campos id  e version são obrigatórios
	if err == nil {
		t.Errorf("Deverá retornar erro, campos id  e version são obrigatórios")
	}
	fmt.Println("App atualizado com sucesso: ", result)

}
