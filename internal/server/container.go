package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/bencoronard/demo-go-crud-api/internal/config"
	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Start() {
	fx.New(
		fx.Provide(
			config.NewProperties,
			config.NewLogger,
			config.NewDB,
			config.NewJwtVerifier,
			config.NewAuthHeaderResolver,
			resource.NewResourceRepoImpl,
			resource.NewResourceServiceImpl,
			resource.NewResourceHandler,
			config.NewRouter,
		),
		fx.Decorate(),
		fx.Invoke(
			startServer,
		),
	).Run()
}

func startServer(lc fx.Lifecycle, r *echo.Echo, log *zap.Logger, props *config.Properties) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info(fmt.Sprintf("Process ID: %d on %s", os.Getpid(), props.Env.App.Environment))
			go func() {
				if err := r.Start(fmt.Sprintf(":%d", props.Env.App.ListenPort)); err != nil && err != http.ErrServerClosed {
					log.Error("HTTP server failed to start", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Shutting down HTTP server...")
			return r.Shutdown(ctx)
		},
	})
}
