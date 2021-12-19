package app_test

import (
	"fmt"
	"testing"

	"github.com/jfonseca85/aws-sdk-dynamodb-expression/app"
)

func Test_CreateApp_Sucess(t *testing.T) {
	//Cenário
	args := map[string]string{
		"id":          "1",
		"cluster":     "itau-service-mesh",
		"squad":       "poseidon",
		"time":        "Os Jedis",
		"quantidades": "23",
		"Teclado":     "23424",
		"Mouse":       "hsjhdkjs",
		"arquivo":     "{}",
		"document":    "QVdTVGVtcGxhdGVGb3JtYXRWZXJzaW9uOiAnMjAxMC0wOS0wOScKUmVzb3VyY2VzOgogIG15U3RhY2s6CiAgICBUeXBlOiBBV1M6OkNsb3VkRm9ybWF0aW9uOjpTdGFjawogICAgUHJvcGVydGllczoKICAgICAgVGVtcGxhdGVVUkw6IGh0dHBzOi8vczMuYW1hem9uYXdzLmNvbS9jbG91ZGZvcm1hdGlvbi10ZW1wbGF0ZXMtdXMtZWFzdC0xL1MzX0J1Y2tldC50ZW1wbGF0ZQogICAgICBUaW1lb3V0SW5NaW51dGVzOiAnNjAnCk91dHB1dHM6CiAgU3RhY2tSZWY6CiAgICBWYWx1ZTogIVJlZiBteVN0YWNrCiAgT3V0cHV0RnJvbU5lc3RlZFN0YWNrOgogICAgVmFsdWU6ICFHZXRBdHQgbXlTdGFjay5PdXRwdXRzLkJ1Y2tldE5hbWU=",
	}
	//Ação
	result, err := app.CreateApp(args)
	if err != nil {
		fmt.Println("Erro durante Criação do App: ", err.Error())
	}

	//Validação
	fmt.Println("Erro durante Criação do App: ", result)
}
