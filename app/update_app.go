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
