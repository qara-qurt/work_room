package handler

import (
	"auth/configs"
	"auth/internal/service"
)

type Handler struct {
	service *service.Service
	cfg     *configs.Config
}

func New(service *service.Service, config *configs.Config) *Handler {
	return &Handler{
		service: service,
		cfg:     config,
	}
}
