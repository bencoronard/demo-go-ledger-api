package server

import (
	"github.com/bencoronard/demo-go-crud-api/internal/config"
	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"go.uber.org/fx"
)

func Start() {
	fx.New(
		fx.Provide(
			config.NewProperties,
			config.NewDB,
			config.NewJwtVerifier,
			config.NewAuthHeaderResolver,
			config.NewRouter,
		),
		fx.Provide(
			resource.NewResourceRepo,
			resource.NewResourceService,
			resource.NewResourceHandler,
		),
		fx.Invoke(
			config.ConfigureLogger,
			config.RegisterMiddlewares,
			config.RegisterRoutes,
		),
		fx.Invoke(config.Start),
	).Run()
}
