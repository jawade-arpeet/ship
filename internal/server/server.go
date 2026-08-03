package server

import (
	"fmt"
	"ship/internal/client"
	"ship/internal/config"
	"ship/internal/handler"
	"ship/internal/middleware"
	"ship/internal/repository"

	"ship/internal/router"
	"ship/internal/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.ServerConfig
	router *gin.Engine
}

func New(
	config *config.ServerConfig,
	client *client.Client,
) *Server {
	repository := repository.New(client)
	service := service.New(repository)
	middleware := middleware.New()
	handler := handler.New(service)
	router := router.New(config.RunEnv, middleware, handler)

	return &Server{config: config, router: router}
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%d", s.config.Port)
	return s.router.Run(addr)
}
