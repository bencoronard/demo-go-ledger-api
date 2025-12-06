package config

import (
	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(h *resource.ResourceHandler) *chi.Mux {
	r := chi.NewRouter()
	registerMiddlewares(r)
	registerRoutes(r, h)
	return r
}

func registerMiddlewares(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}

func registerRoutes(r *chi.Mux, h *resource.ResourceHandler) {
	r.Get("/", h.ListResources)
	r.Get("/", h.RetrieveResource)
	r.Post("/", h.CreateResource)
	r.Put("/", h.UpdateResource)
	r.Delete("/", h.DeleteResource)
}
