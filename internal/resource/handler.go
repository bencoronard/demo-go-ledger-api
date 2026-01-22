package resource

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

func (h *ResourceHandler) RetrieveResource(c echo.Context) error {
	claims, err := h.r.ExtractClaims(c.Request())
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)
	if err != nil {
		return err
	}

	ent, err := h.s.retrieveResource(c.Request().Context(), uint(id), claims)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, toDTO(ent))
}

func (h *ResourceHandler) ListResources(c echo.Context) error {
	claims, err := h.r.ExtractClaims(c.Request())
	if err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	size, _ := strconv.Atoi(c.QueryParam("size"))

	p := dto.NewPageable(page, size)

	sortParams := c.QueryParams()["sort"]
	for _, s := range sortParams {
		parts := strings.Split(s, ",")
		prop := parts[0]
		dir := dto.ASC
		if len(parts) > 1 {
			upperDir := dto.Direction(strings.ToUpper(parts[1]))
			if upperDir == dto.DESC {
				dir = dto.DESC
			}
		}
		if prop != "" {
			p = p.WithSort(prop, dir)
		}
	}

	slice, err := h.s.listResources(c.Request().Context(), p, claims)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.Map(slice, toDTO))
}

func (h *ResourceHandler) CreateResource(c echo.Context) error {
	claims, err := h.r.ExtractClaims(c.Request())
	if err != nil {
		return err
	}

	var body resourceDTO
	if err := c.Bind(&body); err != nil {
		return err
	}

	id, err := h.s.createResource(c.Request().Context(), toEntity(body), claims)
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

	id, err := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)
	if err != nil {
		return err
	}

	var body resourceDTO
	if err := c.Bind(&body); err != nil {
		return err
	}

	if err := h.s.updateResource(c.Request().Context(), uint(id), toEntity(body), claims); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *ResourceHandler) DeleteResource(c echo.Context) error {
	claims, err := h.r.ExtractClaims(c.Request())
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)
	if err != nil {
		return err
	}

	if err := h.s.deleteResource(c.Request().Context(), uint(id), claims); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
