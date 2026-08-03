package handler

import "ship/internal/service"

type Handler struct {
	Health *HealthHandler
}

func New(service *service.Service) *Handler {
	return &Handler{
		Health: newHealthHandler(),
	}
}
