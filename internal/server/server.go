package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"yourproject/config"
	"yourproject/internal/handler"
	"yourproject/internal/service"
	"yourproject/pkg/logger"
)

type Server struct {
	srv    *http.Server
	logger *logger.Logger
}

func New(cfg config.ServerConfig, svc *service.Service, logger *logger.Logger) *Server {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Initialize handlers
	h := handler.New(svc, logger)

	// Routes
	r.Get("/health", h.HealthCheck)

	// API routes
	r.Route("/api/v1", func(r chi.Router) {
		// Add your API routes here
		r.Get("/example", h.GetExample)
		r.Post("/example", h.CreateExample)
	})

	// Create server
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler:      r,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	return &Server{
		srv:    srv,
		logger: logger,
	}
}

func (s *Server) Start() error {
	s.logger.Info("Starting server on", "addr", s.srv.Addr)
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) {
	s.logger.Info("Shutting down server")
	if err := s.srv.Shutdown(ctx); err != nil {
		s.logger.Error("Server shutdown error", "error", err)
	}
}
