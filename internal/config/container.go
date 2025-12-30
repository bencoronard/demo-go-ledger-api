package config

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Container interface {
	Start(lc fx.Lifecycle, sd fx.Shutdowner)
}

type containerImpl struct {
	p *Properties
	l *slog.Logger
	r *echo.Echo
}

func NewContainer(r *echo.Echo, l *slog.Logger, p *Properties) Container {
	return &containerImpl{p: p, l: l, r: r}
}

func (c *containerImpl) Start(lc fx.Lifecycle, sd fx.Shutdowner) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			c.l.Info(fmt.Sprintf("Process ID: %d on %s", os.Getpid(), c.p.Env.App.Environment))

			errChan := make(chan error, 1)

			go func() {
				if err := c.r.Start(fmt.Sprintf(":%d", c.p.Env.App.ListenPort)); err != nil && err != http.ErrServerClosed {
					errChan <- err
				}
			}()

			select {
			case err := <-errChan:
				return err
			case <-time.After(100 * time.Millisecond):
				go func() {
					if err := <-errChan; err != nil {
						c.l.Error(err.Error())
						sd.Shutdown()
					}
				}()
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		},
		OnStop: func(ctx context.Context) error {
			c.l.Info("Shutting down HTTP server...")
			return c.r.Shutdown(ctx)
		},
	})
}
