package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"go.uber.org/fx"
)

type Properties struct {
	Env    EnvCfg
	Secret SecretCfg
}

type EnvCfg struct {
	App   AppCfg
	Vault VaultCfg
	OTEL  OTELCfg
	CP    CPCfg
}

type SecretCfg struct {
	DB DBCfg
}

func NewProperties(lc fx.Lifecycle) (*Properties, error) {
	var props Properties

	if err := env.Parse(&props.Env); err != nil {
		return nil, fmt.Errorf("failed to parse env config: %w", err)
	}

	// lc.Append(fx.Hook{
	// 	OnStart: func(ctx context.Context) error {
	// 		if err := fetchVaultSecrets(ctx, &props.Secret, props.Env.Vault); err != nil {
	// 			return fmt.Errorf("failed to fetch vault secrets: %w", err)
	// 		}
	// 		return nil
	// 	},
	// })

	if err := env.Parse(&props.Secret); err != nil {
		return nil, fmt.Errorf("failed to parse secret config: %w", err)
	}

	return &props, nil
}

// func fetchVaultSecrets(ctx context.Context, s *SecretCfg, cfg VaultCfg) error {
// 	client, err := vault.New(
// 		vault.WithAddress(cfg.URI),
// 		vault.WithRequestTimeout(30*time.Second),
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	if err := client.SetToken(cfg.Token); err != nil {
// 		return err
// 	}

// 	sec, err := client.Secrets.KvV1Read(ctx, "secret")
// 	if err != nil {
// 		return err
// 	}

// 	host, ok := sec.Data["pg.host"].(string)
// 	if !ok {
// 		return errors.New("unable to retrieve PG host")
// 	}

// 	port, ok := sec.Data["pg.port"].(int)
// 	if !ok {
// 		return errors.New("unable to retrieve PG port")
// 	}

// 	s.DB.Host = host
// 	s.DB.Port = port

// 	return nil
// }

type AppCfg struct {
	ListenPort    int    `env:"APP_LISTEN_PORT"`
	Environment   string `env:"APP_ENVIRONMENT"`
	PublicKeyPath string `env:"APP_PUBLIC_KEY_PATH"`
}

type VaultCfg struct {
	URI   string `env:"VAULT_URI"`
	Token string `env:"VAULT_TOKEN"`
}

type OTELCfg struct {
	MetricsEndpoint           string  `env:"OTEL_COL_METRICS_ENDPOINT"`
	TracesEndpoint            string  `env:"OTEL_COL_TRACES_ENDPOINT"`
	LogsEndpoint              string  `env:"OTEL_COL_LOGS_ENDPOINT"`
	MetricsSamplingFreqInMin  string  `env:"OTEL_METRICS_SAMPLING_FREQ_IN_MIN"`
	TracesSamplingProbability float64 `env:"OTEL_TRACES_SAMPLING_PROBABILITY"`
}

type CPCfg struct {
	ConnectionPoolCap         int `env:"CRUD_API_DB_CP_CAP"`
	ConnectionPoolIdleMin     int `env:"CRUD_API_DB_CP_IDLE_MIN"`
	ConnectionPoolIdleTimeout int `env:"CRUD_API_DB_CP_IDLE_TIMEOUT"`
	ConnectionTimeout         int `env:"CRUD_API_DB_CP_CONN_TIMEOUT"`
	ConnectionTTL             int `env:"CRUD_API_DB_CP_CONN_TTL"`
}

type DBCfg struct {
	Host   string `env:"PG_HOST"`
	Port   int    `env:"PG_PORT"`
	DBName string `env:"PG_DBNAME"`
	User   string `env:"PG_USER"`
	Pass   string `env:"PG_PASS"`
}
