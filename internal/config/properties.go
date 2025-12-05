package config

import (
	"strings"

	"github.com/spf13/viper"
)

func ReadProperties() (*Properties, error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var props Properties
	if err := viper.Unmarshal(&props); err != nil {
		return nil, err
	}

	return &props, nil
}

type Properties struct {
	Env    EnvCfg
	Secret SecretCfg
}

type EnvCfg struct {
	Host string `mapstructure:"hostname"`
}

type SecretCfg struct {
	DB DBCfg `mapstructure:"db"`
}

type DBCfg struct {
	Host string `mapstructure:"hostname"`
	Port uint16 `mapstructure:"port"`
}
