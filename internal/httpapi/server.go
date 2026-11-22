// Package httpapi contains transport concerns: routing, HTTP responses, and
// request-scoped metadata. Business logic must not depend on this package.
package httpapi

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

// Server owns the public HTTP handler.
type Server struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Server {
	return &Server{logger: logger}
}

// Handler returns the application's route tree. Go 1.22's ServeMux supports
// method-aware routes without requiring a router dependency.
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", s.health)
	mux.HandleFunc("GET /readyz", s.ready)
	return s.logging(mux)
}

func (s *Server) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) ready(w http.ResponseWriter, _ *http.Request) {
	// In later stages this checks live dependencies (Postgres, Redis and S3).
	// Keeping it distinct from liveness prevents a deadlocked dependency from
	// receiving new traffic while still allowing the process to be inspected.
	writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}

func (s *Server) logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		next.ServeHTTP(w, r)
		s.logger.Info("request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(started).String(),
		)
	})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
