package config

import (
	"net"
	"net/http"
	"strconv"

	web "github.com/bencoronard/go-starter-web/api"
	"github.com/labstack/echo/v5"
	"go.uber.org/fx"
)

type server struct {
	router *echo.Echo
	srv    *http.Server
}

type serverParams struct {
	fx.In
	Router *echo.Echo
	Config web.ServerConfig
}

func NewServer(p serverParams) (web.Server, error) {
	return &server{
		router: p.Router,
		srv: &http.Server{
			Addr:              net.JoinHostPort(p.Config.Host, strconv.Itoa(p.Config.Port)),
			ReadTimeout:       p.Config.ReadTimeout,
			ReadHeaderTimeout: p.Config.ReadHeaderTimeout,
			WriteTimeout:      p.Config.WriteTimeout,
			IdleTimeout:       p.Config.IdleTimeout,
			MaxHeaderBytes:    p.Config.MaxHeaderBytes,
		},
	}, nil
}

func (s *server) Instance() *http.Server {
	return s.srv
}

func (s *server) Configure() error {
	panic("unimplemented")
}
