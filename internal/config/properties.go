package config

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	actuator "github.com/bencoronard/go-starter-actuator/api"
	cache "github.com/bencoronard/go-starter-cache/api"
	jwt "github.com/bencoronard/go-starter-jwt/api"
	rdb "github.com/bencoronard/go-starter-rdb/api"
	vault "github.com/bencoronard/go-starter-vault/api"
	web "github.com/bencoronard/go-starter-web/api"
	"github.com/caarlos0/env/v11"
	"go.uber.org/fx"
)

type envConfig struct {
	ServiceID string `env:"SERVICE_ID"`

	VaultSecretPath           []string `env:"VAULT_SECRET_PATH"`
	VaultSecretReadTimeoutSec int      `env:"VAULT_SECRET_READ_TIMEOUT_SEC"`

	VaultAuthTimeoutSec                     int `env:"VAULT_AUTH_TIMEOUT_SEC"`
	VaultAuthRetryBackoffInitialIntervalSec int `env:"VAULT_AUTH_RETRY_BACKOFF_INTERVAL_MIN_SEC"`
	VaultAuthRetryBackoffMultiplier         int `env:"VAULT_AUTH_RETRY_BACKOFF_MULTIPLIER"`
	VaultAuthRetryBackoffMaxIntervalSec     int `env:"VAULT_AUTH_RETRY_BACKOFF_INTERVAL_MAX_SEC"`

	ActuatorHealthCheckIntervalSec int    `env:"ACTUATOR_HEALTHCHECK_INTERVAL_SEC"`
	ActuatorHealthCheckTimeoutSec  int    `env:"ACTUATOR_HEALTHCHECK_TIMEOUT_SEC"`
	ActuatorServerHost             string `env:"ACTUATOR_SERVER_HOST"`
	ActuatorServerPort             int    `env:"ACTUATOR_SERVER_PORT"`

	RDBMaxOpenConn        int `env:"RDB_MAX_OPEN_CONN"`
	RDBMaxIdleConn        int `env:"RDB_MAX_IDLE_CONN"`
	RDBConnTTLSec         int `env:"RDB_CONN_TTL_SEC"`
	RDBConnIdleTimeoutSec int `env:"RDB_CONN_IDLE_TIMEOUT_SEC"`

	WebServerHost                 string `env:"SERVER_HOST"`
	WebServerPort                 int    `env:"SERVER_PORT"`
	WebServerReadTimeoutSec       int    `env:"SERVER_HTTP_READ_TIMEOUT_SEC"`
	WebServerReadHeaderTimeoutSec int    `env:"SERVER_HTTP_READ_HEADER_TIMEOUT_SEC"`
	WebServerWriteTimeoutSec      int    `env:"SERVER_HTTP_WRITE_TIMEOUT_SEC"`
	WebServerIdleTimeoutSec       int    `env:"SERVER_HTTP_IDLE_TIMEOUT_SEC"`
	WebServerMaxHeaderBytes       int    `env:"SERVER_HTTP_MAX_HEADER_BYTES"`
	WebServerEnableAccessLog      bool   `env:"SERVER_ENABLE_ACCESS_LOG"`
}

type secretConfig struct {
	PrivateKeyPEM string `mapstructure:"key.private"`

	PGHost string `mapstructure:"pg.host"`
	PGPort int    `mapstructure:"pg.port"`
	PGUser string `mapstructure:"pg.user"`
	PGPass string `mapstructure:"pg.pass"`
	PGDB   string `mapstructure:"pg.db"`

	RedisHost     string `mapstructure:"redis.host"`
	RedisPort     int    `mapstructure:"redis.port"`
	RedisPass     string `mapstructure:"redis.pass"`
	RedisDB       int    `mapstructure:"redis.db"`
	RedisProtocol int    `mapstructure:"redis.protocol"`
}

type properties struct {
	fx.Out
	Vault          vault.Config
	Actuator       actuator.ActuatorConfig
	ActuatorServer actuator.ServerConfig
	PGDriver       rdb.PgDriverConfig
	RDB            rdb.ClientConfig
	Cache          cache.RedisClientConfig
	JwtIssuer      jwt.AsymmetricIssuerConfig
	Router         web.EchoRouterConfig
	Server         web.ServerConfig
}

func NewProperties(vc vault.Client) (properties, error) {
	var envCfg envConfig
	if err := env.Parse(&envCfg); err != nil {
		return properties{}, fmt.Errorf("failed to parse config from env: %w", err)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(envCfg.VaultSecretReadTimeoutSec)*time.Second,
	)
	defer cancel()

	var secret secretConfig
	for _, path := range envCfg.VaultSecretPath {
		if err := vc.ReadSecret(ctx, path, &secret); err != nil {
			return properties{}, fmt.Errorf("failed to read secret from path %s: %w", path, err)
		}
	}

	block, _ := pem.Decode([]byte(secret.PrivateKeyPEM))
	if block == nil {
		return properties{}, errors.New("failed to parse private key")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return properties{}, err
	}
	pk, ok := key.(*rsa.PrivateKey)
	if !ok {
		return properties{}, errors.New("not an RSA private key")
	}

	return properties{
		Vault: vault.Config{
			AuthTimeout:                     time.Duration(envCfg.VaultAuthTimeoutSec) * time.Second,
			AuthRetryBackoffInitialInterval: time.Duration(envCfg.VaultAuthRetryBackoffInitialIntervalSec) * time.Second,
			AuthRetryBackoffMaxInterval:     time.Duration(envCfg.VaultAuthRetryBackoffMaxIntervalSec) * time.Second,
			AuthRetryBackoffMultiplier:      envCfg.VaultAuthRetryBackoffMultiplier,
		},
		Actuator: actuator.ActuatorConfig{
			HealthCheckInterval: time.Duration(envCfg.ActuatorHealthCheckIntervalSec) * time.Second,
			HealthCheckTimeout:  time.Duration(envCfg.ActuatorHealthCheckTimeoutSec) * time.Second,
		},
		ActuatorServer: actuator.ServerConfig{
			Host: envCfg.ActuatorServerHost,
			Port: envCfg.ActuatorServerPort,
		},
		PGDriver: rdb.PgDriverConfig{
			Host:     secret.PGHost,
			Port:     secret.PGPort,
			User:     secret.PGUser,
			Password: secret.PGPass,
			DBName:   secret.PGDB,
		},
		RDB: rdb.ClientConfig{
			MaxOpenConns: envCfg.RDBMaxOpenConn,
			MaxIdleConns: envCfg.RDBMaxIdleConn,
			ConnTTL:      time.Duration(envCfg.RDBConnTTLSec) * time.Second,
			IdleTimeout:  time.Duration(envCfg.RDBConnIdleTimeoutSec) * time.Second,
		},
		Cache: cache.RedisClientConfig{
			Host:     secret.RedisHost,
			Port:     secret.RedisPort,
			Password: secret.RedisPass,
			DB:       secret.RedisDB,
			Protocol: secret.RedisProtocol,
		},
		JwtIssuer: jwt.AsymmetricIssuerConfig{
			Issuer: envCfg.ServiceID,
			Key:    pk,
		},
		Router: web.EchoRouterConfig{
			EnableAccessLog: envCfg.WebServerEnableAccessLog,
		},
		Server: web.ServerConfig{
			Host:              envCfg.WebServerHost,
			Port:              envCfg.WebServerPort,
			ReadTimeout:       time.Duration(envCfg.WebServerReadTimeoutSec) * time.Second,
			ReadHeaderTimeout: time.Duration(envCfg.WebServerReadHeaderTimeoutSec) * time.Second,
			WriteTimeout:      time.Duration(envCfg.WebServerWriteTimeoutSec) * time.Second,
			IdleTimeout:       time.Duration(envCfg.WebServerIdleTimeoutSec) * time.Second,
			MaxHeaderBytes:    envCfg.WebServerMaxHeaderBytes,
		},
	}, nil
}
