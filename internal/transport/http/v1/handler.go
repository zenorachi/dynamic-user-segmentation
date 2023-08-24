package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zenorachi/dynamic-user-segmentation/internal/service"
	"github.com/zenorachi/dynamic-user-segmentation/pkg/auth"
)

type Handler struct {
	services     *service.Services
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Services, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initUserRoutes(v1)
	}
}
