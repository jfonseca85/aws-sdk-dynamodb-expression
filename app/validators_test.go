package app_test

import (
	"testing"

	"github.com/jfonseca85/aws-sdk-dynamodb-expression/app"
)

func Test_Validators_Params_Correct(t *testing.T) {
	//Cenário: Cria o args com parametros corretos

	args := map[string]string{
		"id":       "1",
		"document": "QVdTVGVtcGxhdGVGb3JtYXRWZXJzaW9uOiAnMjAxMC0wOS0wOScKUmVzb3VyY2VzOgogIG15U3RhY2s6CiAgICBUeXBlOiBBV1M6OkNsb3VkRm9ybWF0aW9uOjpTdGFjawogICAgUHJvcGVydGllczoKICAgICAgVGVtcGxhdGVVUkw6IGh0dHBzOi8vczMuYW1hem9uYXdzLmNvbS9jbG91ZGZvcm1hdGlvbi10ZW1wbGF0ZXMtdXMtZWFzdC0xL1MzX0J1Y2tldC50ZW1wbGF0ZQogICAgICBUaW1lb3V0SW5NaW51dGVzOiAnNjAnCk91dHB1dHM6CiAgU3RhY2tSZWY6CiAgICBWYWx1ZTogIVJlZiBteVN0YWNrCiAgT3V0cHV0RnJvbU5lc3RlZFN0YWNrOgogICAgVmFsdWU6ICFHZXRBdHQgbXlTdGFjay5PdXRwdXRzLkJ1Y2tldE5hbWU=",
	}

	//Ação: Realização a validação do args
	err := app.ValidateParams(args, app.CreateAppParams())

	//Validação: Não deve retornar erros campos obrigatórios preenchidos ( id e document )
	if err != nil {
		t.Errorf("Erro na validação do Parametros: %v/n", err.Error())
	}
}

func Test_Validators_Id_Missing(t *testing.T) {
	//Cenário: Cria args sem o id (campo obrigatório )

	args := map[string]string{
		"version":  "v4",
		"document": "QVdTVGVtcGxhdGVGb3JtYXRWZXJzaW9uOiAnMjAxMC0wOS0wOScKUmVzb3VyY2VzOgogIG15U3RhY2s6CiAgICBUeXBlOiBBV1M6OkNsb3VkRm9ybWF0aW9uOjpTdGFjawogICAgUHJvcGVydGllczoKICAgICAgVGVtcGxhdGVVUkw6IGh0dHBzOi8vczMuYW1hem9uYXdzLmNvbS9jbG91ZGZvcm1hdGlvbi10ZW1wbGF0ZXMtdXMtZWFzdC0xL1MzX0J1Y2tldC50ZW1wbGF0ZQogICAgICBUaW1lb3V0SW5NaW51dGVzOiAnNjAnCk91dHB1dHM6CiAgU3RhY2tSZWY6CiAgICBWYWx1ZTogIVJlZiBteVN0YWNrCiAgT3V0cHV0RnJvbU5lc3RlZFN0YWNrOgogICAgVmFsdWU6ICFHZXRBdHQgbXlTdGFjay5PdXRwdXRzLkJ1Y2tldE5hbWU=",
	}

	//Ação: Realiza a validação dos args
	err := app.ValidateParams(args, app.CreateAppParams())

	//Validação: A validação deverá retornar erro, está faltando o id (obrigatório)
	if err == nil {
		t.Errorf("A validação deverá retornar erro, está faltando o ID: %v/n", err.Error())
	}
}

//	Neste test não passamos o campo version,
//	Validação deverá preencher criar o campo e preencher com latest
//TODO(Jorge Luis): Lógica desabilitada na operação de createApp

/*
func Test_Validators_Version_Missing(t *testing.T) {

	//Cenário
	args := map[string]string{
		"id":       "12",
		"document": "QVdTVGVtcGxhdGVGb3JtYXRWZXJzaW9uOiAnMjAxMC0wOS0wOScKUmVzb3VyY2VzOgogIG15U3RhY2s6CiAgICBUeXBlOiBBV1M6OkNsb3VkRm9ybWF0aW9uOjpTdGFjawogICAgUHJvcGVydGllczoKICAgICAgVGVtcGxhdGVVUkw6IGh0dHBzOi8vczMuYW1hem9uYXdzLmNvbS9jbG91ZGZvcm1hdGlvbi10ZW1wbGF0ZXMtdXMtZWFzdC0xL1MzX0J1Y2tldC50ZW1wbGF0ZQogICAgICBUaW1lb3V0SW5NaW51dGVzOiAnNjAnCk91dHB1dHM6CiAgU3RhY2tSZWY6CiAgICBWYWx1ZTogIVJlZiBteVN0YWNrCiAgT3V0cHV0RnJvbU5lc3RlZFN0YWNrOgogICAgVmFsdWU6ICFHZXRBdHQgbXlTdGFjay5PdXRwdXRzLkJ1Y2tldE5hbWU=",
	}
	expected := "latest"

	//Ação
	err := app.ValidateParams(args, app.CreateAppParams())
	result := args["version"]

	//Resultado
	if result != expected {
		t.Errorf("Erro na validação dos parametro a versão deveria ser latest porem o valor retornado foi : %v/n", result)
	}

	if err != nil {
		t.Errorf("Erro interno na validação : %v/n", err.Error())
	}

}
*/

/*
	Neste teste não passamos o campo document (obrigatório),
	A validação  deverá retornar um erro informando que este campo é obrigatório.
*/
func Test_Validators_Document_Missing(t *testing.T) {
	//Cenário: Cria args sem o document (campo obrigatório )
	args := map[string]string{
		"id": "37",
	}

	//Validação: A validação deverá retornar erro, está faltando o document (obrigatório)
	err := app.ValidateParams(args, app.CreateAppParams())
	//Validação: A validação deverá retornar erro, está faltando o document (obrigatório)
	if err == nil {
		t.Errorf("Erro na validação o campo document é obrigatório")
	}
}

func Test_Validators_Params_Additional(t *testing.T) {

	//Cenário: Cria args com campos adicionais
	args := map[string]string{
		"id":       "37",
		"version":  "v4",
		"document": "QVdTVGVtcGxhdGVGb3JtYXRWZXJzaW9uOiAnMjAxMC0wOS0wOScKUmVzb3VyY2VzOgogIG15U3RhY2s6CiAgICBUeXBlOiBBV1M6OkNsb3VkRm9ybWF0aW9uOjpTdGFjawogICAgUHJvcGVydGllczoKICAgICAgVGVtcGxhdGVVUkw6IGh0dHBzOi8vczMuYW1hem9uYXdzLmNvbS9jbG91ZGZvcm1hdGlvbi10ZW1wbGF0ZXMtdXMtZWFzdC0xL1MzX0J1Y2tldC50ZW1wbGF0ZQogICAgICBUaW1lb3V0SW5NaW51dGVzOiAnNjAnCk91dHB1dHM6CiAgU3RhY2tSZWY6CiAgICBWYWx1ZTogIVJlZiBteVN0YWNrCiAgT3V0cHV0RnJvbU5lc3RlZFN0YWNrOgogICAgVmFsdWU6ICFHZXRBdHQgbXlTdGFjay5PdXRwdXRzLkJ1Y2tldE5hbWU=",
		"cell":     "<path of cell>",
		"xpto":     "value_xpto",
	}

	//Ação: Valida os args
	err := app.ValidateParams(args, app.CreateAppParams())

	//Validação: Deverá permitir inserir campos adicionais
	if err != nil {
		t.Errorf("Deverá permitir inserir campos adicionais")
	}
}

func Test_Validators_Params_Empty(t *testing.T) {
	//Cenário: Cria args vazio, campos id e document são obrigatórios
	args := map[string]string{}

	//Ação; Validar os args
	err := app.ValidateParams(args, app.CreateAppParams())

	//Validação:Deverá retornar erro, campor ID e document são obrigatórios
	if err == nil {
		t.Errorf("Deverá retornar erro, campor ID e document são obrigatórios")
	}
}

func Test_Validators_Params_Null(t *testing.T) {
	//Cenário: Args igual a Null

	//Ação: Valida o args igual a Null
	err := app.ValidateParams(nil, app.CreateAppParams())

	//Validação: Deverá retornar erro, campor ID e document são obrigatórios
	if err == nil {
		t.Errorf("Deverá retornar erro, campor ID e document são obrigatórios")
	}
}
