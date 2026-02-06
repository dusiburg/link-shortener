package http

import "example.com/m/internal/handlers"

func (s *Server) routes() {
	api := s.engine

	api.GET("/ping", handlers.Ping)
	api.GET("/l/*link", s.linksHandler.CreateLink)
	s.engine.GET("/r/:id", s.linksHandler.Redirect)
}
