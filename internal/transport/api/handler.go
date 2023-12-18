package api

import (
	"github.com/gin-gonic/gin"
	"tests/internal/services"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() {
	router := gin.Default()
	api := router.Group("/api")
	{
		h.InitEntRoutes(api)
	}
}
