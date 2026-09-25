package server

import (
	"context"
	"fmt"
	"ship/internal/client"
	"ship/internal/config"
	"ship/internal/handler"
	"ship/internal/middleware"
	"ship/internal/repository"
	"ship/internal/router"
	"ship/internal/service"

	"github.com/labstack/echo/v5"
)

type Server struct {
	config *config.ServerConfig
	router *echo.Echo
}

func New() (*Server, error) {
	cfg := config.GetServerConfig()

	ctx := context.Background()

	clt, err := client.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	repo := repository.New(clt)
	svc := service.New(repo)
	mw := middleware.New()
	hdlr := handler.New(svc)
	rtr := router.New(mw, hdlr)

	return &Server{
		config: cfg,
		router: rtr,
	}, nil
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.config.Port)
	return s.router.Start(addr)
}
