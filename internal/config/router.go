package config

import (
	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	return echo.New()
}

func RegisterMiddlewares(e *echo.Echo) {
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
}

func RegisterRoutes(e *echo.Echo, h *resource.ResourceHandler) {
	e.GET("/", h.ListResources)
	e.GET("/", h.RetrieveResource)
	e.POST("/", h.CreateResource)
	e.PUT("/", h.UpdateResource)
	e.DELETE("/", h.DeleteResource)
}
