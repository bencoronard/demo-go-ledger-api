package resource

import (
	"github.com/labstack/echo/v4"
)

type ResourceHandler struct {
	s resourceService
}

func NewResourceHandler(s resourceService) *ResourceHandler {
	return &ResourceHandler{s: s}
}

func (h *ResourceHandler) ListResources(c echo.Context) error {
	panic("unimplemented")
}

func (h *ResourceHandler) CreateResource(c echo.Context) error {
	panic("unimplemented")
}

func (h *ResourceHandler) DeleteResource(c echo.Context) error {
	panic("unimplemented")
}

func (h *ResourceHandler) RetrieveResource(c echo.Context) error {
	panic("unimplemented")
}

func (h *ResourceHandler) UpdateResource(c echo.Context) error {
	panic("unimplemented")
}
