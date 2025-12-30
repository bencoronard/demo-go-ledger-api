package config

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

func NewRouter(lc fx.Lifecycle, h *resource.ResourceHandler, p *Properties) *http.Server {
	e := echo.New()
	registerMiddlewares(e)
	registerRoutes(e, h)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", p.Env.App.ListenPort),
		Handler: e,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

func registerMiddlewares(e *echo.Echo) {
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
}

func registerRoutes(e *echo.Echo, h *resource.ResourceHandler) {
	e.GET("/", h.ListResources)
	e.GET("/", h.RetrieveResource)
	e.POST("/", h.CreateResource)
	e.PUT("/", h.UpdateResource)
	e.DELETE("/", h.DeleteResource)
}
