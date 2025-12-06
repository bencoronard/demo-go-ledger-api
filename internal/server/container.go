package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/bencoronard/demo-go-crud-api/internal/config"
	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Start() {
	fx.New(
		fx.Provide(
			config.ReadProperties,
			config.NewLogger,
			config.NewDB,
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

func startServer(lc fx.Lifecycle, r *chi.Mux, logger *zap.Logger, props *config.Properties) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info(fmt.Sprintf("Process ID: %d on %s", os.Getpid(), props.Env.Host))
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Error("HTTP server failed to start", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down HTTP server...")
			return srv.Shutdown(ctx)
		},
	})
}
