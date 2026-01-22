package config

import (
	"context"
	"fmt"
	"net"
	"net/http"

	xhttp "github.com/bencoronard/demo-go-common-libs/http"
	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type router struct {
	port int
	e    *echo.Echo
	h    *resource.ResourceHandler
}

func NewRouter(h *resource.ResourceHandler, p *Properties) xhttp.Router {
	return &router{
		port: p.Env.App.ListenPort,
		e:    echo.New(),
		h:    h,
	}
}

func (r *router) ListeningPort() int {
	return r.port
}

func (r *router) Listen(port int) (net.Listener, error) {
	addr := fmt.Sprintf(":%d", port)
	return net.Listen("tcp", addr)
}

func (r *router) Serve(l net.Listener) error {
	return r.e.Server.Serve(l)
}

func (r *router) Shutdown(ctx context.Context) error {
	return r.e.Shutdown(ctx)
}

func (r *router) RegisterMiddlewares() {
	r.e.Use(middleware.RequestLogger())
	r.e.Use(middleware.Recover())
}

func (r *router) RegisterRoutes() {
	api := r.e.Group("/api")
	api.GET("/resources/:id", r.h.RetrieveResource)
	api.GET("/resources", r.h.ListResources)
	api.POST("/resources", r.h.CreateResource)
	api.PUT("/resources/:id", r.h.UpdateResource)
	api.DELETE("/resources/:id", r.h.DeleteResource)

	act := r.e.Group("/actuator")
	act.GET("/health", func(c echo.Context) error { return c.JSON(http.StatusOK, map[string]string{"status": "up"}) })
	act.GET("/readiness", func(c echo.Context) error { return c.JSON(http.StatusOK, map[string]string{"status": "ready"}) })
}
