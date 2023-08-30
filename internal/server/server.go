package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zenorachi/dynamic-user-segmentation/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func New(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port),
			Handler:      handler,
			ReadTimeout:  cfg.HTTP.ReadTimeout,
			WriteTimeout: cfg.HTTP.WriteTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
