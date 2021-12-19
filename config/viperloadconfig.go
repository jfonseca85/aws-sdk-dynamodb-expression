package config

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/viper"
)

type Config struct {
	Viper     *viper.Viper
	AWSConfig aws.Config
}

func NewConfig(ctx context.Context) (*Config, error) {
	viper.AddConfigPath(".")  // to work on dev and production envs
	viper.AddConfigPath("./") // to work on dev and production envs
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	awsconfig, err := awsloadconfig(ctx, viper.GetViper())
	if err != nil {
		return nil, err
	}

	return &Config{
		Viper:     viper.GetViper(),
		AWSConfig: awsconfig,
	}, nil
}
