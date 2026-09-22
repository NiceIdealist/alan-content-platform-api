package server

import (
	"github.com/gin-gonic/gin"

	"github.com/alan-content-platform/alan-content-platform-api/internal/handler"
)

// Server wraps the Gin engine for the API.
type Server struct {
	engine *gin.Engine
}

// New builds a Server with routes registered.
func New() *Server {
	engine := gin.Default()
	registerRoutes(engine)

	return &Server{engine: engine}
}

// Run starts the HTTP server on the given address.
func (s *Server) Run(addr string) error {
	return s.engine.Run(addr)
}

func registerRoutes(engine *gin.Engine) {
	engine.GET("/healthz", handler.Health)
}
