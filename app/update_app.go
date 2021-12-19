package app

import (
	"context"
	"fmt"

	"github.com/jfonseca85/aws-sdk-dynamodb-expression/configlocal"
)

func UpdateApp(ctx context.Context, cfg *configlocal.Viperloadconfig, args map[string]string) (*Model, error) {
	fmt.Println("Invoke UpdateApp")
	err := ValidateParams(args, updateAppParams())
	if err != nil {
		return nil, err
	}
	return UpdateAppMiddleware(ctx, cfg, args)
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
