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

func startServer(lc fx.Lifecycle, r *echo.Echo, logger *zap.Logger, props *config.Properties) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info(fmt.Sprintf("Process ID: %d on %s", os.Getpid(), props.Env.Host))
			go func() {
				if err := r.Start(":8080"); err != nil && err != http.ErrServerClosed {
					logger.Error("HTTP server failed to start", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down HTTP server...")
			return r.Shutdown(ctx)
		},
	})
}
