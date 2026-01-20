package main

import (
	"github.com/bencoronard/demo-go-common-libs/http"
	"github.com/bencoronard/demo-go-common-libs/orm"
	"github.com/bencoronard/demo-go-crud-api/internal/config"
	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.NewProperties,
			config.NewDB,
			orm.NewTransactionManager,
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
			http.Router.RegisterMiddlewares,
			http.Router.RegisterRoutes,
		),
		fx.Invoke(http.Start),
	).Run()
}
