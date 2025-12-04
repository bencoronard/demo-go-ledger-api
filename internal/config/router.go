package config

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	registerMiddlewares(r)
	registerRoutes(r)
	return r
}

func registerMiddlewares(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}

func registerRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!\n"))
	})
}
