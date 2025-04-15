package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HTTPServer struct {
	server *http.Server
}

func NewHTTPServer(address string, router *chi.Mux) *HTTPServer {
	return &HTTPServer{
		server: &http.Server{
			Addr:    address,
			Handler: router,
		},
	}
}

func (s *HTTPServer) Start() error {
	return s.server.ListenAndServe()
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
