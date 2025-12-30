package config

import (
	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(h *resource.ResourceHandler) *echo.Echo {
	r := echo.New()
	r.HideBanner = true
	r.HidePort = true
	return r
}

func RegisterMiddlewares(r *echo.Echo) {
	r.Use(middleware.RequestLogger())
	r.Use(middleware.Recover())
}

func RegisterRoutes(r *echo.Echo, h *resource.ResourceHandler) {
	r.GET("/", h.ListResources)
	r.GET("/", h.RetrieveResource)
	r.POST("/", h.CreateResource)
	r.PUT("/", h.UpdateResource)
	r.DELETE("/", h.DeleteResource)
}
