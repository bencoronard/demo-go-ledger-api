package config

import (
	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

func NewProperties() (*Properties, error) {
	v := viper.New()
	v.AutomaticEnv()

	keys := []string{
		"APP_LISTEN_PORT",
		"APP_ENVIRONMENT",
		"APP_PUBLIC_KEY_PATH",
		"VAULT_URI", "VAULT_TOKEN",
		"OTEL_COL_METRICS_ENDPOINT",
		"OTEL_COL_TRACES_ENDPOINT",
		"OTEL_COL_LOGS_ENDPOINT",
		"OTEL_METRICS_SAMPLING_FREQ_IN_MIN",
		"OTEL_TRACES_SAMPLING_PROBABILITY",
		"CRUD_API_DB_CP_CAP",
		"CRUD_API_DB_CP_IDLE_MIN",
		"CRUD_API_DB_CP_IDLE_TIMEOUT",
		"CRUD_API_DB_CP_CONN_TIMEOUT",
		"CRUD_API_DB_CP_CONN_TTL",
	}

	for _, key := range keys {
		v.BindEnv(key)
	}

	var props Properties
	err := v.Unmarshal(&props, func(c *mapstructure.DecoderConfig) {
		c.TagName = "mapstructure"
	})

	if err != nil {
		return nil, err
	}

	return &props, nil
}

type Properties struct {
	Env EnvCfg `mapstructure:",squash"`
}

type EnvCfg struct {
	App   AppCfg   `mapstructure:",squash"`
	Vault VaultCfg `mapstructure:",squash"`
	OTEL  OTELCfg  `mapstructure:",squash"`
	DB    DBCfg    `mapstructure:",squash"`
}

type AppCfg struct {
	ListenPort    int    `mapstructure:"APP_LISTEN_PORT"`
	Environment   string `mapstructure:"APP_ENVIRONMENT"`
	PublicKeyPath string `mapstructure:"APP_PUBLIC_KEY_PATH"`
}

type VaultCfg struct {
	URI   string `mapstructure:"VAULT_URI"`
	Token string `mapstructure:"VAULT_TOKEN"`
}

type OTELCfg struct {
	MetricsEndpoint           string  `mapstructure:"OTEL_COL_METRICS_ENDPOINT"`
	TracesEndpoint            string  `mapstructure:"OTEL_COL_TRACES_ENDPOINT"`
	LogsEndpoint              string  `mapstructure:"OTEL_COL_LOGS_ENDPOINT"`
	MetricsSamplingFreqInMin  string  `mapstructure:"OTEL_METRICS_SAMPLING_FREQ_IN_MIN"`
	TracesSamplingProbability float64 `mapstructure:"OTEL_TRACES_SAMPLING_PROBABILITY"`
}

type DBCfg struct {
	ConnectionPoolCap         int `mapstructure:"CRUD_API_DB_CP_CAP"`
	ConnectionPoolIdleMin     int `mapstructure:"CRUD_API_DB_CP_IDLE_MIN"`
	ConnectionPoolIdleTimeout int `mapstructure:"CRUD_API_DB_CP_IDLE_TIMEOUT"`
	ConnectionTimeout         int `mapstructure:"CRUD_API_DB_CP_CONN_TIMEOUT"`
	ConnectionTTL             int `mapstructure:"CRUD_API_DB_CP_CONN_TTL"`
}
