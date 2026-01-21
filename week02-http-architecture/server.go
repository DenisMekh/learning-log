package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	server *http.Server
}

func NewServer(router http.Handler) *Server {
	return &Server{
		&http.Server{
			Addr:    ":8080",
			Handler: router,
		},
	}
}

func NewRouter(handler Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(RequestIDMiddleware)
	r.Use(LoggingMiddleware)
	r.Post("/health", handler.Health)
	r.Post("/ping", handler.Ping)
	r.Post("/long", handler.SomeLong)
	return r
}

func (s *Server) Start() error {
	log.Println("Starting server...")
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Shutting down server...")
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
