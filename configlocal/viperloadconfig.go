package configlocal

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/viper"
)

type Viperloadconfig struct {
	Viper     *viper.Viper
	AWSClient aws.Config
}

func NewConfig(ctx context.Context) (*Viperloadconfig, error) {
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

	return &Viperloadconfig{
		Viper:     viper.GetViper(),
		AWSClient: awsconfig,
	}, nil
}
