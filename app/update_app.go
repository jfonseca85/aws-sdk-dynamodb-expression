//Esse adaptador de driver deve ser capaz de transformar uma solicitação http em uma chamada de serviço.
//Ter todos os componentes separados uns dos outros nos dá a vantagem de implementá-los e
//Testá-los isoladamente ou podemos facilmente paralelizar o trabalho com a ajuda de outros membros da equipe.
package app

import (
	"fmt"
)

func UpdateApp(args map[string]string) (*Model, error) {
	fmt.Println("Invoke UpdateApp")
	err := ValidateParams(args, updateAppParams())
	if err != nil {
		return nil, err
	}
	return UpdateAppMiddleware(args)
}

func updateAppParams() []*Param {
	return []*Param{
		{
			Name:     argId,
			Type:     "string",
			Required: true,
		},
		{
			Name:     argVersion,
			Type:     "string",
			Required: true,
		},
	}
}
