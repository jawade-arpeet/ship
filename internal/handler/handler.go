package handler

import "ship/internal/service"

type Handler struct {
	Health *HealthHandler
}

func New(svc *service.Service) *Handler {
	return &Handler{
		Health: newHealthHandler(),
	}
}
