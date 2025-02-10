package server

import (
	"avito-shop-service/internal/config"
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func New(h http.Handler, cfg config.Server) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         cfg.Address,
			Handler:      h,
			IdleTimeout:  cfg.IdleTimeout,
			ReadTimeout:  cfg.Timeout,
			WriteTimeout: cfg.Timeout,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
