package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/bencoronard/demo-go-crud-api/internal/config"
	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"github.com/labstack/echo/v4"
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
		fx.Invoke(start),
	).Run()
}

func start(lc fx.Lifecycle, sd fx.Shutdowner, e *echo.Echo, p *config.Properties) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			slog.Info(fmt.Sprintf("Process ID: %d on %s", os.Getpid(), p.Env.App.Environment))

			errChan := make(chan error, 1)

			go func() {
				if err := e.Start(fmt.Sprintf(":%d", p.Env.App.ListenPort)); err != nil && err != http.ErrServerClosed {
					errChan <- err
				}
			}()

			select {
			case <-ctx.Done():
				return ctx.Err()
			case err := <-errChan:
				return err
			case <-time.After(100 * time.Millisecond):
				go func() {
					if err := <-errChan; err != nil {
						slog.Error(err.Error())
						sd.Shutdown()
					}
				}()
				return nil
			}
		},
		OnStop: func(ctx context.Context) error {
			slog.Info("Shutting down HTTP server...")
			return e.Shutdown(ctx)
		},
	})
}
