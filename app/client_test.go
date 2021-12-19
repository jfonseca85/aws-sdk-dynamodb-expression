package app_test

import (
	"context"

	"github.com/jfonseca85/aws-sdk-dynamodb-expression/config"
)

func NewClientLocal() (context.Context, *config.Config, error) {
	ctx := context.TODO()
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		return nil, nil, err
	}
	return ctx, cfg, nil
}
