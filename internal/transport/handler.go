package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swagFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zenorachi/dynamic-user-segmentation/internal/config"
	"github.com/zenorachi/dynamic-user-segmentation/internal/service"
	v1 "github.com/zenorachi/dynamic-user-segmentation/internal/transport/http/v1"
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

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(config.New().GIN.Mode)

	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.GET("/docs/*any", ginSwagger.WrapHandler(swagFiles.Handler))

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(r *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.tokenManager)
	api := r.Group("/api")
	{
		handlerV1.Init(api)
	}
}
