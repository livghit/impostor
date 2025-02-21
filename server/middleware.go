package server

import "github.com/go-chi/chi/v5/middleware"

func (s *Server) LoadMiddleware() {
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
}
