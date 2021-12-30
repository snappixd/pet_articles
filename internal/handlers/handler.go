package handler

import (
	"articles/internal/config"
	"articles/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("ui/html/*")

	h.initArticlesRoutes(router)

	return router
}
