package config

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"go.uber.org/fx"
)

type Container interface {
	Run() error
}

type container struct {
	srv *http.Server
}

func NewContainer(h *http.Server) Container {
	return &container{srv: h}
}

func Start(lc fx.Lifecycle, sd fx.Shutdowner, c Container) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			errChan := make(chan error, 1)

			go func() {
				if err := c.Run(); err != nil && err != http.ErrServerClosed {
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
	})
}

func (c *container) Run() error {
	return c.srv.ListenAndServe()
}
