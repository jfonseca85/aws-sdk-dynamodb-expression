package app

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jfonseca85/aws-sdk-dynamodb-expression/config"
)

func UpdateApp(ctx context.Context, cfg *config.Config, args map[string]string) (string, error) {
	fmt.Println("Invoke UpdateApp")
	err := ValidateParams(args, UpdateAppParams())
	if err != nil {
		return "", err
	}
	result, _ := UpdateAppMiddleware(ctx, cfg, args)

	ret, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(ret), nil
}

func UpdateAppParams() []*Param {
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
