package app_test

import (
	"testing"

	"github.com/jfonseca85/aws-sdk-dynamodb-expression/app"
)

/*
	Neste teste passamos todos os campos obrigatórios (id e document )
	A validação deverá retornar sem erros :)
*/
func Test_Validators_Params_Correct(t *testing.T) {
	//Cenário

	args := map[string]string{
		"id":       "1",
		"version":  "v4",
		"document": "QVdTVGVtcGxhdGVGb3JtYXRWZXJzaW9uOiAnMjAxMC0wOS0wOScKUmVzb3VyY2VzOgogIG15U3RhY2s6CiAgICBUeXBlOiBBV1M6OkNsb3VkRm9ybWF0aW9uOjpTdGFjawogICAgUHJvcGVydGllczoKICAgICAgVGVtcGxhdGVVUkw6IGh0dHBzOi8vczMuYW1hem9uYXdzLmNvbS9jbG91ZGZvcm1hdGlvbi10ZW1wbGF0ZXMtdXMtZWFzdC0xL1MzX0J1Y2tldC50ZW1wbGF0ZQogICAgICBUaW1lb3V0SW5NaW51dGVzOiAnNjAnCk91dHB1dHM6CiAgU3RhY2tSZWY6CiAgICBWYWx1ZTogIVJlZiBteVN0YWNrCiAgT3V0cHV0RnJvbU5lc3RlZFN0YWNrOgogICAgVmFsdWU6ICFHZXRBdHQgbXlTdGFjay5PdXRwdXRzLkJ1Y2tldE5hbWU=",
	}

	//Ação
	err := app.ValidateParams(args, app.CreateAppParams())
	//Resultado
	if err != nil {
		t.Errorf("Erro na validação do Parametros: %v/n", err.Error())
	}
}

/*
	Neste teste não passamos o campo id (obrigatório),
	A validação  deverá retornar um erro com os campo faltantes.
*/
func Test_Validators_Id_Missing(t *testing.T) {
	//Cenário

	args := map[string]string{
		"version":  "v4",
		"document": "QVdTVGVtcGxhdGVGb3JtYXRWZXJzaW9uOiAnMjAxMC0wOS0wOScKUmVzb3VyY2VzOgogIG15U3RhY2s6CiAgICBUeXBlOiBBV1M6OkNsb3VkRm9ybWF0aW9uOjpTdGFjawogICAgUHJvcGVydGllczoKICAgICAgVGVtcGxhdGVVUkw6IGh0dHBzOi8vczMuYW1hem9uYXdzLmNvbS9jbG91ZGZvcm1hdGlvbi10ZW1wbGF0ZXMtdXMtZWFzdC0xL1MzX0J1Y2tldC50ZW1wbGF0ZQogICAgICBUaW1lb3V0SW5NaW51dGVzOiAnNjAnCk91dHB1dHM6CiAgU3RhY2tSZWY6CiAgICBWYWx1ZTogIVJlZiBteVN0YWNrCiAgT3V0cHV0RnJvbU5lc3RlZFN0YWNrOgogICAgVmFsdWU6ICFHZXRBdHQgbXlTdGFjay5PdXRwdXRzLkJ1Y2tldE5hbWU=",
	}

	//Ação
	err := app.ValidateParams(args, app.CreateAppParams())
	//Resultado
	if err == nil {
		t.Errorf("A validação deverá retornar erro, está faltando o ID: %v/n", err.Error())
	}
}

/*
	Neste test não passamos o campo version,
	Validação deverá preencher criar o campo e preencher com latest
*/
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

/*
	Neste teste não passamos o campo document (obrigatório),
	A validação  deverá retornar um erro informando que este campo é obrigatório.
*/
func Test_Validators_Document_Missing(t *testing.T) {
	//Cenário

	args := map[string]string{
		"id":      "37",
		"version": "v4",
	}

	//Ação
	err := app.ValidateParams(args, app.CreateAppParams())
	//Resultado
	if err == nil {
		t.Errorf("Erro na validação o campo document é obrigatório")
	}
}
