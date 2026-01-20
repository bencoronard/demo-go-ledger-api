package resource

import (
	"fmt"
	"net/http"

	"github.com/bencoronard/demo-go-common-libs/dto"
	xhttp "github.com/bencoronard/demo-go-common-libs/http"
	"github.com/labstack/echo/v4"
)

type ResourceHandler struct {
	r xhttp.AuthHeaderResolver
	s resourceService
}

func NewResourceHandler(r xhttp.AuthHeaderResolver, s resourceService) *ResourceHandler {
	return &ResourceHandler{r: r, s: s}
}

func (h *ResourceHandler) ListResources(c echo.Context) error {
	claims, err := h.r.ExtractClaims(c.Request())
	if err != nil {
		return err
	}

	res, err := h.s.listResources(c.Request().Context(), dto.Pageable{}, claims)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *ResourceHandler) RetrieveResource(c echo.Context) error {
	claims, err := h.r.ExtractClaims(c.Request())
	if err != nil {
		return err
	}

	res, err := h.s.retrieveResource(c.Request().Context(), 0, claims)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *ResourceHandler) CreateResource(c echo.Context) error {
	claims, err := h.r.ExtractClaims(c.Request())
	if err != nil {
		return err
	}

	id, err := h.s.createResource(c.Request().Context(), nil, claims)
	if err != nil {
		return err
	}

	c.Response().Header().Set(echo.HeaderLocation, fmt.Sprintf("/api/resources/%d", id))
	return c.NoContent(http.StatusCreated)
}

func (h *ResourceHandler) UpdateResource(c echo.Context) error {
	claims, err := h.r.ExtractClaims(c.Request())
	if err != nil {
		return err
	}

	if err := h.s.updateResource(c.Request().Context(), 0, nil, claims); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *ResourceHandler) DeleteResource(c echo.Context) error {
	claims, err := h.r.ExtractClaims(c.Request())
	if err != nil {
		return err
	}

	if err := h.s.deleteResource(c.Request().Context(), 0, claims); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
