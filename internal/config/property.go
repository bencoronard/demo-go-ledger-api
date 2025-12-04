package config

import "github.com/spf13/viper"

type AppProps struct {
	Host string
}

func ReadProperties() *AppProps {
	viper.AutomaticEnv()

	return &AppProps{
		Host: viper.GetString("HOSTNAME"),
	}
}
