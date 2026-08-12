package main

import (
	"github.com/bencoronard/demo-go-rest-api/internal/config"
	actuator "github.com/bencoronard/go-starter-actuator/api"
	cache "github.com/bencoronard/go-starter-cache/api"
	jwt "github.com/bencoronard/go-starter-jwt/api"
	otel "github.com/bencoronard/go-starter-otel/api"
	rdb "github.com/bencoronard/go-starter-rdb/api"
	validation "github.com/bencoronard/go-starter-validation/api"
	vault "github.com/bencoronard/go-starter-vault/api"
	web "github.com/bencoronard/go-starter-web/api"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		// Vault
		fx.Provide(
			vault.NewVaultClient,
			vault.NewClient,
			vault.NewAppRoleAuth,
			vault.NewWatcher,
		),
		// Cache
		fx.Provide(
			cache.NewRedisClient,
			cache.NewRedisCounter,
			cache.NewRedisKVBlobStore,
			cache.NewRedisKVFieldStore,
		),
		// Actuator
		fx.Provide(
			actuator.NewActuator,
			actuator.NewServer,
		),
		// JWT
		fx.Provide(
			jwt.NewAsymmetricIssuer,
			jwt.NewAsymmetricVerifier,
		),
		// OpenTelemetry
		fx.Provide(
			otel.NewResource,
			otel.NewPropagator,
			otel.NewTracerProvider,
			otel.NewMeterProvider,
			otel.NewLoggerProvider,
			otel.NewLogger,
		),
		// Validation
		fx.Provide(
			validation.NewValidator,
		),
		// Web
		fx.Provide(
			web.NewAuthHeaderResolver,
			web.NewEchoRouter,
			config.NewServer,
		),
		// Relational DB
		fx.Provide(
			rdb.NewPgDriver,
			rdb.NewClient,
			rdb.NewHealthChecker,
		),
		// Properties
		fx.Provide(
			config.NewProperties,
		),
		// Entry Point
		fx.Invoke(
			web.StartServer,
		),
	).Run()
}
