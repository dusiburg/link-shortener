package http

import (
	"fmt"
	"log"
	"net/http"

	"example.com/m/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	address string

	engine *gin.Engine

	linksHandler *handlers.LinksHandler
}

func NewServer(address string, lh *handlers.LinksHandler) *Server {
	engine := gin.New()

	engine.Use(gin.Recovery())

	return &Server{
		address:      address,
		engine:       engine,
		linksHandler: lh,
	}
}

func (s *Server) Start() error {
	s.routes()

	if err := s.engine.Run(s.address); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error occured while starting the server: %v", err)
		return err
	}

	fmt.Println("server started: &v", s.address)
	return nil
}
