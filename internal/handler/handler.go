package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/quanergyO/avito_assingment/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.SignUp)
		auth.POST("sign-in", h.SignIn)
	}

	return router
}
